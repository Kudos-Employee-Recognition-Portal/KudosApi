package handlers

import (
	"cloud.google.com/go/storage"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/models"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Note: r, w used for request and response objects respectively by emerging convention in golang apis.
func GetUsers(db *sql.DB) http.Handler {
	// Return the handler as a closure over the database object.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			// Write errors explicitly. Could be changed to http response text later.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encoding error is explicitly ignored as data structure is verified in model method.
		_ = json.NewEncoder(w).Encode(users)
	})
}

func GetAdmins(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllAdmins(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(users)
	})
}

func GetManagers(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		managers, err := models.GetAllManagers(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(managers)
	})
}

func GetUser(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := models.User{Email: vars["email"]}
		err := user.GetUserByEmail(db)
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
		_ = json.NewEncoder(w).Encode(user)
	})
}

func GetAdmin(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		admin := models.User{ID: id}
		err = admin.GetUserById(db)
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
		_ = json.NewEncoder(w).Encode(admin)
	})
}

func GetManager(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		manager := models.User{ID: id}
		err = manager.GetManagerById(db)
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
		_ = json.NewEncoder(w).Encode(manager)
	})
}

func GetManagerAwards(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		manager := models.User{ID: id}
		awards, err := manager.GetManagerAwards(db)
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

func CreateAdmin(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var admin models.User
		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		admin.Type = "admin"

		err = admin.CreateAdmin(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(admin)
	})
}

func UpdateAdmin(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var admin models.User
		err = json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		admin.ID = id
		err = admin.UpdateUserInfo(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = admin.GetUserById(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(admin)
	})
}

func CreateManager(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var manager models.User
		err := json.NewDecoder(r.Body).Decode(&manager)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		manager.Type = "manager"

		err = manager.CreateManager(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(manager)
	})
}

func UpdateManager(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var manager models.User
		err = json.NewDecoder(r.Body).Decode(&manager)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		manager.ID = id
		err = manager.UpdateUserInfo(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = manager.UpdateManagerInfo(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = manager.GetUserById(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(manager)
	})
}

func SetManagerSignature(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		manager := models.User{ID: id}

		// Strange but true, idiomatic use a left shift operator to set maxMemory to 20MB
		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusRequestEntityTooLarge)
		}
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Connect to the Google Cloud Storage bucket where signatures are stored.
		ctx := context.Background()
		client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
		// TODO: switch to inferred credentials when deployed, maybe.
		//client, err := storage.NewClient(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		bkt := client.Bucket(os.Getenv("G_BUCKET_NAME"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		ctx, cancel := context.WithTimeout(ctx, time.Second*60)
		defer cancel()
		objWriter := bkt.Object("user" + strconv.Itoa(id) + "signature").NewWriter(ctx)
		_, err = io.Copy(objWriter, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Trying out the other idiomatic error checking structure here.
		if err = objWriter.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		manager.SigURL.String = "https://storage.cloud.google.com/kudosapi.appspot.com/user" + strconv.Itoa(id) + "signature"
		err = manager.UpdateManagerSignature(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(manager)
	})
}

func DeleteUser(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user := models.User{ID: id}
		err = user.DeleteUser(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	})
}
