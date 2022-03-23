package main

import (
	aart "asciiweb/pkg"
	"html/template"
	"log"
	"net/http"
)

type TextToShow struct {
	Datext string
}

var NEWTEXT = TextToShow{""}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art", asciiart)

	log.Println("Starting Server")
	err := http.ListenAndServe(":7630", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404: Page Not Found"))
		// log.Println(http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFiles("templates/index.tmpl")
	// should have another if statement to check if exists
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, NEWTEXT)
	// error when cannot do the execution of the template
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Bad Request"))
		log.Println(http.StatusBadRequest)
		return
	}
}

// we need to POST to /ascii-art
func asciiart(w http.ResponseWriter, r *http.Request) {
	// if method is not POST then http status 400 for badrequest
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Bad Request"))
		log.Println(http.StatusBadRequest)
		http.Redirect(w, r, "/", http.StatusBadRequest) // Not sure again
		return
	}

	text := r.FormValue("textinput")
	for _, i := range text {
		if i < 0 || i > 127 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400: Bad Request"))
			log.Println(http.StatusBadRequest)
			return
		}
	}
	banner := r.FormValue("banner")
	// handle http status for when banner is not selected
	if banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Bad Request"))
		log.Println(http.StatusBadRequest)
		return
	}

	log.Printf("Received user input: %q", text)
	log.Printf("Banner has been selected: %q", banner)

	asciiconverted := aart.GenerateArt(text, banner)
	log.Println("Successfully Converted")

	NEWTEXT.Datext = asciiconverted

	http.Redirect(w, r, "/", http.StatusFound) // nik : not sure about this status
}
