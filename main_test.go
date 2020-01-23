package main

import(
	"github.com/joho/godotenv"
	"net/http"
	"net/http/httptest"
	"os"
	"log"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	// Load environment variables from .env file with godotenv package.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	a = App{}
	a.Init(os.Getenv("API_DB_USERNAME"), os.Getenv("API_DB_PASSWORD"), os.Getenv("API_DB_NAME"))

	assertTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func assertTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL
)`

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	if body := res.Body.String(); body != "[]" {
		t.Errorf("Expected empty array, but got %s.", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	resRec := httptest.NewRecorder()
	a.Router.ServeHTTP(resRec, req)

	return resRec
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected: %d; Actual: %d\n", expected, actual)
	}
}