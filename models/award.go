package models

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"database/sql"
	"encoding/base64"
	"github.com/go-sql-driver/mysql"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/api/option"
	"image/png"
	_ "io"
	"io/ioutil"
	"log"
	_ "net/http"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"
)

// Award struct reflects the award db entity after being joined to relevant tables, receiving values rather than keys.
type Award struct {
	ID             int            `json:"id"`
	Region         region         `json:"region"`
	Type           string         `json:"type"`
	RecipientName  string         `json:"recipientname"`
	RecipientEmail string         `json:"recipientemail"`
	CreatedBy      User           `json:"createdby"`
	CreatedOn      mysql.NullTime `json:"createdon"`
	// and an accessory field to make queries easier, not sure if best practice.
	QueryDates dateRange `json:"daterange"`
}

// Future: implement graphQL interface.

type region struct {
	ID   int    `json:"regionid"`
	Name string `json:"regionname"`
}

type dateRange struct {
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type Awards []Award
type Regions []region

func GetAllAwards(db *sql.DB) (Awards, error) {
	rows, err := db.Query(
		"SELECT a.id, r.id, r.name, a.type, a.recipientName, a.recipientEmail, a.createdOn, m.user_id, m.firstName, m.lastName " +
			"FROM award a " +
			"JOIN region r ON a.region_id = r.id " +
			"JOIN manager m ON a.createdBy = m.user_id ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards Awards
	for rows.Next() {
		var award Award
		err := rows.Scan(
			&award.ID, &award.Region.ID, &award.Region.Name, &award.Type, &award.RecipientName,
			&award.RecipientEmail, &award.CreatedOn,
			&award.CreatedBy.ID, &award.CreatedBy.FirstName, &award.CreatedBy.LastName)
		if err != nil {
			return nil, err
		}
		awards = append(awards, award)
	}
	return awards, nil
}

func (award *Award) QueryAwards(db *sql.DB) (Awards, error) {
	rows, err := db.Query("SELECT a.id, r.id, r.name, a.type, a.recipientName, a.recipientEmail, a.createdOn, "+
		"m.user_id, m.firstName, m.lastName "+
		"FROM award a "+
		"JOIN region r ON a.region_id = r.id "+
		"JOIN manager m ON a.createdBy = m.user_id "+
		"WHERE (a.createdOn BETWEEN ? AND ?) "+
		"AND (? IS NULL OR a.type LIKE ?) "+
		"AND (? IS NULL OR a.recipientName LIKE ?) "+
		"AND (? IS NULL OR a.recipientEmail LIKE ?) "+
		"AND (? IS NULL OR r.name LIKE ?) ",
		award.QueryDates.StartDate, award.QueryDates.EndDate,
		award.Type, "%"+award.Type+"%",
		award.RecipientName, "%"+award.RecipientName+"%",
		award.RecipientEmail, "%"+award.RecipientEmail+"%",
		award.Region.Name, "%"+award.Region.Name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards Awards
	for rows.Next() {
		var award Award
		err := rows.Scan(
			&award.ID, &award.Region.ID, &award.Region.Name, &award.Type, &award.RecipientName,
			&award.RecipientEmail, &award.CreatedOn,
			&award.CreatedBy.ID, &award.CreatedBy.FirstName, &award.CreatedBy.LastName)
		if err != nil {
			return nil, err
		}
		awards = append(awards, award)
	}
	return awards, nil
}

func (award *Award) GetAward(db *sql.DB) error {
	return db.QueryRow(
		"SELECT a.id, r.id, r.name, a.type, a.recipientName, a.recipientEmail, a.createdOn, "+
			"m.user_id, m.firstName, m.lastName, m.signatureURL "+
			"FROM award a "+
			"JOIN region r ON a.region_id = r.id "+
			"JOIN manager m ON a.createdBy = m.user_id "+
			"WHERE a.id = ?",
		award.ID).Scan(
		&award.ID, &award.Region.ID, &award.Region.Name, &award.Type, &award.RecipientName,
		&award.RecipientEmail, &award.CreatedOn,
		&award.CreatedBy.ID, &award.CreatedBy.FirstName, &award.CreatedBy.LastName,
		&award.CreatedBy.SigURL)
}

func (award *Award) CreateAward(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO award (region_id, type, recipientName, recipientEmail, createdBy) "+
			"VALUES (?, ?, ?, ?, ?)",
		&award.Region.ID, &award.Type, &award.RecipientName,
		&award.RecipientEmail, &award.CreatedBy.ID)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	award.ID = int(insertID)
	return nil
}

func (award *Award) DeleteAward(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM award WHERE id = ?",
		award.ID)
	return err
}

func (award *Award) EmailAward(filename string) error {
	// Try V3 API for attachments.
	email := mail.NewV3Mail()
	// Set sender.
	email.SetFrom(mail.NewEmail("Kudos!", "awardsteam-no-reply@kudosapi.appspotmail.com"))
	// Set content.
	email.AddContent(mail.NewContent("text/html", "<h1>"+award.Type+"<h1><h2>Congratulations "+award.RecipientName+"!!</h2>"))

	// Personalization, add award recipient logic here.
	personalization := mail.NewPersonalization()
	personalization.AddTos(mail.NewEmail(award.RecipientName, award.RecipientEmail))
	personalization.Subject = award.CreatedBy.FirstName + " gave you an award. Great Job!!"
	email.AddPersonalizations(personalization)

	// Process file to attachment.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	fileAttachment := mail.NewAttachment()
	fileAttachment.SetContent(base64.StdEncoding.EncodeToString(data))
	fileAttachment.SetType("application/pdf")
	fileAttachment.SetFilename("kudos.pdf")
	fileAttachment.SetDisposition("attachment")
	fileAttachment.SetContentID("Your Kudos Award")

	// Add attachment to email.
	email.AddAttachment(fileAttachment)

	// Build request object.
	request := sendgrid.GetRequest(os.Getenv("KUDOS_API_SENDGRID"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(email)
	// Ship it!
	_, err = sendgrid.API(request)
	if err != nil {
		return err
	}
	return nil
}

func (award *Award) GetSignatureImage(dname string) (string, error) {
	// Get the image from cloud storage.
	url := award.CreatedBy.SigURL.String
	// If the manager hasn't set a signature, pretend to be Bruce Lee.
	if url == "" {
		return "2846902_2.jpg", nil
	}

	// Open up the signature bucket.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Println("Cloud storage error.")
		return "", err
	}
	// Read out the image blob.
	imageObject, err := client.Bucket(os.Getenv("G_BUCKET_NAME")).Object("user3signature").NewReader(ctx)
	if err != nil {
		return "", err
	}
	// Read the blob to a byte string.
	imageData, err := ioutil.ReadAll(imageObject)
	if err != nil {
		return "", err
	}
	// Decode the byte string as a png image type.
	image, err := png.Decode(bytes.NewReader(imageData))

	// open file to tempdir to save signature.
	signatureFilepath := filepath.Join(dname, "signature.png")
	signatureFile, err := os.Create(signatureFilepath)
	if err != nil {
		return "", err
	}
	defer signatureFile.Close()

	// Encode the image as png and write to temp file.
	err = png.Encode(signatureFile, image)
	if err != nil {
		log.Println("Failed to encode PNG.")
		return "", nil
	}

	// Return the temporary signature file path.
	return signatureFilepath, nil
}

func (award *Award) Tex2Pdf(dname string, signatureFilepath string) (string, error) {
	// Insert relevant award object variables into tex template.
	type awardTemplate struct {
		Title     string
		Region    string
		Date      string
		Recipient string
		Manager   string
		Signature string
	}
	awardFields := awardTemplate{
		Title:     award.Type,
		Region:    award.Region.Name,
		Date:      award.CreatedOn.Time.Format("January 02, 2006"),
		Recipient: award.RecipientName,
		Manager:   award.CreatedBy.FirstName + " " + award.CreatedBy.LastName,
		Signature: signatureFilepath,
	}
	templateOutputFile, err := os.OpenFile("award.tex", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("failed to open output file")
		return "", err
	}
	defer templateOutputFile.Close()
	templateInputFile, err := template.New("award.gotex").Delims("##", "##").ParseFiles("award.gotex")
	if err != nil {
		log.Println("failed to parse")
		return "", err
	}

	err = templateInputFile.Execute(templateOutputFile, awardFields)
	if err != nil {
		log.Println("failed to execute template")
		return "", err
	}

	// Convert tex to pdf and save to the temp directory.
	// TODO: try to reduce os calls: pdflatex -halt-on-error -output-directory dname award.tex
	cmd := exec.Command("pdflatex", "award.tex")
	if err := cmd.Run(); err != nil {
		log.Println("pdflatex conversion failure")
		return "", err
	}
	pdfFilepath := filepath.Join(dname, "award.pdf")
	cmd = exec.Command("mv", "award.pdf", pdfFilepath)
	if err := cmd.Run(); err != nil {
		log.Println("os exec command failure.")
		return "", err
	}

	// Clean up and return:
	return pdfFilepath, nil
}

func GetAllRegions(db *sql.DB) (Regions, error) {
	rows, err := db.Query(
		"SELECT id, name FROM region")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regions Regions
	for rows.Next() {
		var region region
		err := rows.Scan(
			&region.ID, &region.Name)
		if err != nil {
			return nil, err
		}
		regions = append(regions, region)
	}
	return regions, nil
}
