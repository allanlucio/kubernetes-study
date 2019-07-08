package main

import (
	"fmt"
	"net/http"
)
func greeting(nome string) string {
	return "<b>" +nome+ "</b>"
}
func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, greeting("Code.education Rocks!"))
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", nil)
}