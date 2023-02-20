package main

import (
	"log"
	"net/http"
	"text/template"
)

func MainHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	myTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("template error : %s", err.Error())
	}
	myTemplate.Execute(responseWriter, map[string]interface{}{
		"input": "testing hello",
	})
}

func main() {
	http.HandleFunc("/", MainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
