package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ViewData struct {
	Title   string
	Message string
}

type Users struct {
	Username string
	Chatid   int64
	Version  string
	Status   int64
	Purchase int64
	Period   int64
	Signal_1 string
	Signal_2 string
	Signal_3 string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := ViewData{
			Title:   "World Cup",
			Message: "FIFA will never regret it",
		}
		tmpl, _ := template.ParseFiles("templates/Menu2.2.html")
		var users Users
		GetDBhtml(&users)
		tmpl.Execute(w, users)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

func GetDBhtml(users *Users) {
	result := GetDB().Where("username = ?", "PadreAlex").Find(&users)
	if result.Error != nil {
		return
	}
	return
}
