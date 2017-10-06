package main

import (
	"net/http"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[1:]
    http.ServeFile(w, r, title)
}

func signUp(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("address: " + r.Form["address"][0])
    fmt.Println(r.Form)
    http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
    http.HandleFunc("/signup/submit", signUp)
    http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}
