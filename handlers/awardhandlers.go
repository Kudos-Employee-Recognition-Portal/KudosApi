package handlers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/models"
	"github.com/gorilla/mux"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetAwards(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awards, err := models.GetAllAwards(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(awards)
	})
}

func QueryAwards(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var award models.Award
		award.QueryDates.StartDate = r.URL.Query().Get("startdate")
		if award.QueryDates.StartDate == "" {
			award.QueryDates.StartDate = "2000-01-01"
		}
		award.QueryDates.EndDate = r.URL.Query().Get("enddate")
		if award.QueryDates.EndDate == "" {
			award.QueryDates.EndDate = time.Now().String()
		}
		award.Type = r.URL.Query().Get("awardtype")
		award.RecipientName = r.URL.Query().Get("recipientname")
		award.RecipientEmail = r.URL.Query().Get("recipientemail")
		award.Region.Name = r.URL.Query().Get("regionname")
		awards, err := award.QueryAwards(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(awards)
	})
}

func CreateAward(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var award models.Award
		err := json.NewDecoder(r.Body).Decode(&award)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = award.CreateAward(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: Insert function call here to generate LaTeX to PDF to email chain from populated award.
		// tex2pdf

		// Email via Twilio SendGrid API integration.
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
		// TODO: change to PDF when conversion working.
		data, err := ioutil.ReadFile("README.md")
		if err != nil {
			http.Error(w, "Couldn't read file.", http.StatusInternalServerError)
			return
		}
		fileAttachment := mail.NewAttachment()
		fileAttachment.SetContent(base64.StdEncoding.EncodeToString([]byte(data)))
		fileAttachment.SetType("text/plain")
		fileAttachment.SetFilename("certificate.md")
		fileAttachment.SetDisposition("attachment")
		fileAttachment.SetContentID("Test Attachment")

		// Add attachment to email.
		email.AddAttachment(fileAttachment)

		// Build request object.
		request := sendgrid.GetRequest(os.Getenv("KUDOS_API_SENDGRID"), "/v3/mail/send", "https://api.sendgrid.com")
		request.Method = "POST"
		request.Body = mail.GetRequestBody(email)
		// Ship it!
		response, err := sendgrid.API(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			// TODO: remove success dev log.
			log.Println(response.StatusCode)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(award)
	})
}

func GetAward(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		award := models.Award{ID: id}
		err := award.GetAward(db)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(award)
	})
}

func DeleteAward(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		award := models.Award{ID: id}
		err = award.DeleteAward(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	})
}
