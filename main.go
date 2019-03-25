package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
   "github.com/gorilla/mux"
)

var tmpl_ = template.Must(template.ParseGlob("../templates/*.html"))

func processIndex(w http.ResponseWriter, r *http.Request) {
   fmt.Println("/dashboard")
   err := tmpl_.ExecuteTemplate(w, "main", "data")
   if err != nil { panic(err) }
}

func main() {
   r := mux.NewRouter()
   r.HandleFunc("/", processIndex)

   // Register the path "styles" so that the main.html can find the .css files
   r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("../templates/styles/"))))


	fmt.Println("ListenAndServe: http://localhost:80")
   http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", nil))
	// log.Fatal(http.ListenAndServeTLS(":443", "Resources/server.crt", "Resources/server.key", nil))
}

