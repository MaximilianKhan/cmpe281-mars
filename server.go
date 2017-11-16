package main

import (
	"fmt"
	"net/http"
	"github.com/go-redis/redis"
	"html/template"
	"io/ioutil"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Image struct {
	img_name string
	img_addr string
}

type Page struct {
	numOfProducts int
	images []Image
}

func renderPage(w http.ResponseWriter, page string, p *Page){
	t, _ := template.ParseFiles(page)
	t.Execute(w, p)
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	password, _ := client.Get(r.Form["email"][0]).Result()
	if (password == r.Form["pswd"])
	{
		http.Redirect(w, r, "/", http.StatusFound)
		fmt.Println("Password", password)
	}
	else {
		http.Redirect(w, r, "/", http.StatusNoContent)
	}
}

func (pg *Page) getPage(page *Page){
	//connect to MongoDB
	session, err := mgo.Dial("localhost")//aws address
	if err != nil {
			panic(err)
	}
	defer session.Close()
	//querry DB
	c := session.DB("").C("")
	//set page
	return (page)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	page := &Page{numOfProducts: 0, images: []}
	//connect to MongoDB database and pull product items
	page := getPage(page)
	renderPage(w, title, items)
	// http.ServeFile(w, r, title)
}
func signUp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.SetNX(r.Form["username"][0], r.Form["password"][0], 0).Err()
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.HandleFunc("/signup/submit", signUp)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8082", nil)
}