package model

import (
	"database/sql"
	"text/template"

	_ "github.com/lib/pq"
)

type People struct {
	Username   string
	Name       string
	Password   string
	Repassword string
	Permission string
}

type AdminDetails struct {
	Users     []People
	AdminName string
}

type UserDetails struct {
	Name     string
	Username string
}

type ErrorDetails struct {
	EmailError    string
	PasswordError string
	NameError     string
	CommonError   string
	Redirected    bool
}

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=adithyaadhi dbname=adminpanel sslmode=disable")
	if err != nil {
		panic(err)
	}
	//defer DB.Close()
}

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*"))
}
