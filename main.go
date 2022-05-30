package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template
var connStr string
var db *sql.DB

func init() {
	fmt.Printf("Starting application port:8080 \n")

	var err error
	connStr = "host=localhost port=5432 dbname=databasename user=database password=password sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("NO OPEN, err")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	} else {
		fmt.Println("Database connected:")
	}

	tmpl = template.Must(template.ParseGlob("pages/*html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}
func signupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "signup.html", nil)
}
func trailerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "trailer.html", nil)
}

func main() {
	defer db.Close()
	fs := http.FileServer(http.Dir("assets"))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/trailer", trailerHandler)
	http.ListenAndServe(":8080", nil)
}
