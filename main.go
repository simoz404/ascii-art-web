package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
 
	ascii "ascii-web/Functions"
)


func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", submitHandler)

	fmt.Println("http://localhost:8080\nStarting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "error 405 not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	if r.URL.Path != "/" && r.URL.Path != "/submit" {
		
		http.Error(w, "error 404 not found ", http.StatusNotFound)
		return
	}
	t, err2 := template.ParseFiles("Templates/index.html")
	if err2 != nil {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}

	err3 := t.Execute(w, "Submit text")
	if err3 != nil {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "error 404 not found ", http.StatusNotFound)
		return
	}
	text := r.FormValue("input")
	banner := r.FormValue("banner")
	s, err := ascii.RetuenAscii(text, banner)
	if err {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	} else if text == "" {
		http.Error(w, "error 400 bad request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	t2, err1 := template.ParseFiles("Templates/output.html")
	if err1 != nil {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
	t2.ExecuteTemplate(w, "output.html", s)
}
