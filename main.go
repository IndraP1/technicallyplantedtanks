package main

import (
	"net/http"
	"html/template"
  "logger"
	// "log"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))


func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
  log := logger.InitServerLogger()
  defer log.Close()

	logger.Print("Starting TPT site")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

  fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

  mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
