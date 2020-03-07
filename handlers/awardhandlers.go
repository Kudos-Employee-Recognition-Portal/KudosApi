package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/models"
	"github.com/gorilla/mux"
	"io/ioutil"
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		err = award.GetAward(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a temp directory to handle assets.
		dname, err := ioutil.TempDir("", "texpdf")
		if err != nil {
			if err0 := award.DeleteAward(db); err0 != nil {
				http.Error(w, err0.Error(), http.StatusInternalServerError)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Defer removal of the temp dir, texpdf, until function completion.
		// TODO: write a deferred error checker.
		defer os.RemoveAll(dname)

		// Get the awarding manager's signature image from cloud storage and save to the tempdir,
		//	dname, to use as an asset in pdf creation.
		signatureFilepath, err := award.GetSignatureImage(dname)
		if err != nil {
			if err0 := award.DeleteAward(db); err0 != nil {
				http.Error(w, err0.Error(), http.StatusInternalServerError)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the award to tex to pdf, save the pdf to the tempdir specified by dname,
		//	and return the string name of the generated pdf's filepath.
		pdfFilepath, err := award.Tex2Pdf(dname, signatureFilepath)
		if err != nil {
			if err0 := award.DeleteAward(db); err0 != nil {
				http.Error(w, err0.Error(), http.StatusInternalServerError)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Email via Twilio SendGrid API integration.
		// Email the award recipient
		if err = award.EmailAward(pdfFilepath); err != nil {
			if err0 := award.DeleteAward(db); err0 != nil {
				http.Error(w, err0.Error(), http.StatusInternalServerError)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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

func GetRegions(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		regions, err := models.GetAllRegions(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(regions)
	})
}
