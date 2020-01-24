package app

import (
	"../routes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) InitDB(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	// Initialize sql db connection pool and assign to the App struct.
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database.")
	}
}

func (app *App) InitRouter() {
	// Create new gorilla mux router and assign to the App struct.
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/", routes.ApiInfo)

	usersRouter := app.Router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/", routes.GetUsers).Methods("GET")
	usersRouter.HandleFunc("/managers", routes.GetManagers).Methods("GET")
	usersRouter.HandleFunc("/managers", routes.CreateManager).Methods("POST")
	usersRouter.HandleFunc("/managers/{id}", routes.GetManager).Methods("GET")
	usersRouter.HandleFunc("/managers/{id}", routes.UpdateManager).Methods("PUT")
	usersRouter.HandleFunc("/managers/{id}", routes.DeleteManager).Methods("DELETE")
	usersRouter.HandleFunc("/admins", routes.GetAdmins).Methods("GET")
	usersRouter.HandleFunc("/admins", routes.CreateAdmin).Methods("POST")
	usersRouter.HandleFunc("/admins/{id}", routes.GetAdmin).Methods("GET")
	usersRouter.HandleFunc("/admins/{id}", routes.UpdateAdmin).Methods("PUT")
	usersRouter.HandleFunc("/admins/{id}", routes.DeleteAdmin).Methods("DELETE")

}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

