package app

import (
	"database/sql"
	"fmt"
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Wrapping the router and database objects in a struct has several benefits. For this application, primarily the
//	ability to define functions that take the struct as a receiver allowing scoped access to its components.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Database initialization assumes an sql database using a mysql driver. Connection string components
//	are passed as arguments and composed on initialization, allowing more flexibility in
func (app *App) InitDB(user, password, host, port, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	// Initialize sql db connection pool and assign to the App struct.
	var err error
	// The golang sql module opens the database as a connection pool by default.
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database.")
	}
	// However, sql.Open() does not verify the connection. Therefore, best practice is to ping the connection pool before proceeding.
	err = app.DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database ping successful.")
	}
}

// Initializing the router defines the mux module options and primary routing paths using mux subrouters.
func (app *App) InitRouter() {
	// Create new gorilla mux router and assign to the App struct's Router property.
	app.Router = mux.NewRouter()

	// TODO: Restrict requests to server domain after deployment. This will secure the api against abusive or malformed requests.

	// Handle routes by either directly passing a handler function or pointing to a subrouter directing function.
	app.Router.Handle("/", ApiInfo(app.DB))
	routes.UsersRouter(app.Router.PathPrefix("/users").Subrouter(), app.DB)
	routes.AwardsRouter(app.Router.PathPrefix("/awards").Subrouter(), app.DB)
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
