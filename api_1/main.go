package main

import (
	"fmt"
	"html"
	"html/template"

	"math/rand"
	"net/http"
	"strconv"
	"strings"
	_ "time"
)

var Numerals = map[int]string{
	10: "X",
	9:  "IX",
	8:  "VIII",
	7:  "VII",
	6:  "VI",
	5:  "V",
	4:  "IV",
	3:  "III",
	2:  "II",
	1:  "I",
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/randomint", randomInt)
	router.HandleFunc("/randomfloat", randomfloat)
	// http package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// If request is GET with correct syntax
		if strings.TrimSpace(urlPathElements[1]) != "romannumeral" {
			w.Write([]byte("Wrong Path"))
			return
		}
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		fmt.Println(number)
		if number == 0 || number > 10 {
			// If resource is not in the list, send Not Found status
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%q",
				html.EscapeString(Numerals[number]))
		}

		// For all other requests, tell that Client sent a bad request
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("400 - Bad request"))

	})

	// Create a server and run it on 8000 port
	// s := &http.Server{
	// 	Addr:           ":8000",
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	//s.ListenAndServe()
	//s.ListenAndServe()
	http.ListenAndServe(":8080", router)
}
func randomInt(w http.ResponseWriter, r *http.Request) {
	file := template.Must(template.ParseFiles("index.html"))
	file.Execute(w, rand.Intn(100))

}
func randomfloat(w http.ResponseWriter, r *http.Request) {
	file := template.Must(template.ParseFiles("index.html"))
	file.Execute(w, rand.Float64())

}
