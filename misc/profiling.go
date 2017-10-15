package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type add struct {
	Sum int
}
var templates = template.Must(template.ParseFiles("template.html"))

func handleStructAdd(w http.ResponseWriter, r *http.Request) {
	var html bytes.Buffer
	// extract value from FormValue
	first, second := r.FormValue("first"), r.FormValue("second")
	// converting to int
	one, err := strconv.Atoi(first)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	two, err := strconv.Atoi(second)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// dumping to map
	m := struct{ a, b int }{one, two}
	// dumping back to add struct
	structSum := add{Sum: m.a + m.b}

	// html template
	err = templates.Execute(&html, structSum)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html.String()))

}

func main() {
	http.HandleFunc("/struct", handleStructAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
