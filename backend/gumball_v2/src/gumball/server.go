/*
	Products API in Go (Version 2)
	Uses MongoDB and RabbitMQ 
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "localhost"
var mongodb_database = "cmpe281"
var mongodb_collection = "products"


// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/gumball", gumballHandler(formatter)).Methods("GET")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API Gumball Machine Handler
func gumballHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Products:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}


/*

	-- RabbitMQ Setup

	http://localhost:8080

	-- RabbitMQ Create Queue:  

		Queue Name: gumball
		Durable:	no

	-- Gumball MongoDB Create Database

		Database Name: cmpe281
		Collection Name: products

  	-- Gumball MongoDB Collection (Create Document) --

	db.products.insert(
	    { 
	      Id: 1,
	      Count: NumberInt(202),
	      ModelNumber: 'M102988',
	      SerialNumber: '1234998871109' 
	    }
	) ;

    -- products MongoDB Collection - Find Gumball Document --

    db.products.find( { Id: 1 } ) ;


    -- products MongoDB Collection - Update Gumball Document --

    db.products.update( 
        { Dd: 1 }, 
        { $set : { Count : NumberInt(10) } },
        { multi : false } 
    )

    -- products Delete Documents

    db.products.remove({})

 */
