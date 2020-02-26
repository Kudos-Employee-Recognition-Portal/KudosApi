package models

import (
	"database/sql"
	"encoding/base64"
	"github.com/go-sql-driver/mysql"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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
	email.SetFrom(mail.NewEmail("Kudos!", "awardsteam@kudosapi.appspotmail.com"))
	// Set content.
	email.AddContent(mail.NewContent("text/html", "<h2>Congratulations!!</h2>"))

	// Personalization, add award recipient logic here.
	personalization := mail.NewPersonalization()
	personalization.AddTos(mail.NewEmail("McDude", "mcdadem@oregonstate.edu"))
	personalization.Subject = "Someone gave you an award. Great Job!!"
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
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	} else {
		// TODO: remove success dev log.
		log.Println(response.StatusCode)
	}
	return nil
}

func (award *Award) GetSignatureImage(dname string) (string, error) {
	// Get the image from cloud storage.
	url := award.CreatedBy.SigURL.String
	if url == "" {
		return "2846902_2.jpg", nil
	}
	// Could also get image via cloud storage, but this takes advantage of
	//	having the image url available.
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Save it to the tempdir.
	signatureFilepath := filepath.Join(dname, "signature.jpeg")
	signatureFile, err := os.Create(signatureFilepath)
	if err != nil {
		return "", err
	}
	defer signatureFile.Close()
	_, err = io.Copy(signatureFile, response.Body)
	if err != nil {
		return "", err
	}

	// Return the temporary signature file path.
	return signatureFilepath, nil
}

func (award *Award) Tex2Pdf(dname string, signatureFilepath string) (string, error) {
	// Insert relevant award object variables into tex template.

	// Convert tex to pdf and save to the temp directory.

	cmd := exec.Command("pdflatex", "test.tex")
	if err := cmd.Run(); err != nil {
		return "", err
	}
	pdfFilepath := filepath.Join(dname, "award.pdf")
	cmd = exec.Command("mv", "test.pdf", pdfFilepath)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Clean up and return:
	return pdfFilepath, nil
}
