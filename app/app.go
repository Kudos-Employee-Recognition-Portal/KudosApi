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

func (app *App) InitDB(user, password, host, port, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	// Initialize sql db connection pool and assign to the App struct.
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database.")
	}
	err = app.DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database ping successful.")
	}

	// show available tables in the connected database.
	//rows, _ := app.DB.Query("SHOW TABLES")
	//defer rows.Close()
	//for rows.Next() {
	//	var name string
	//	_ = rows.Scan(&name)
	//	log.Println(name)
	//}
}

func (app *App) InitRouter() {
	// Create new gorilla mux router and assign to the App struct's Router property.
	app.Router = mux.NewRouter()

	// TODO: Restrict requests to server domain after deployment.

	// Handle routes by either directly passing a handler function or pointing to a subrouter directing function.
	app.Router.Handle("/", ApiInfo(app.DB))
	routes.UsersRouter(app.Router.PathPrefix("/users").Subrouter(), app.DB)
	routes.AwardsRouter(app.Router.PathPrefix("/awards").Subrouter(), app.DB)
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
