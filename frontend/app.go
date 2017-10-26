package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

type ucred struct {
	username string
	password string
}

func NewClient(u *ucred) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)
	// Output: PONG <nil>
	err := client.Set(u.username, u.password, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(u.username).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(u.username, val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}

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
