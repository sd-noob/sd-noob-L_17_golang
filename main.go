package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first golang webpage")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Name - Sourav Das \nElectrical and Electronic Engineer(EEE)")
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Email ID: sdhriday77@gmail.com \n Phone no: 017XXXXXXXX")
}
