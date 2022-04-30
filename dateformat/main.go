package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/form", home2)          // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	file, _ := template.ParseFiles("index.html")

	file.Execute(w, nil)

}

func home2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "index.html", 302)
	date := r.URL.Query().Get("date")
	fmt.Println(date)
	w.Write([]byte(date))

}
