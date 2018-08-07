package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// global reference
// fine for this toy example because the web server owns this entire package
var collection *mongo.Collection

func init() {

	// TODO investigate how this interacts with a docker service
	//      seems to be the main problem in containerizing this
	//      with docker-compose
	client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal("creating client: ", err)
	}
	log.Printf("client: %+v\n", client)

	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal("connecting to client: ", err)
	}

	collection = client.Database("Rewards").Collection("customers")
	log.Printf("collection: %+v\n", collection)
}

func main() {

	// TODO use packr (https://github.com/gobuffalo/packr) to wrap
	//      up this entire app in a Go binary so that deployment
	//      is as easy as dropping an executable in an environment
	//      and configuring a process manager

	r := mux.NewRouter()
	r.HandleFunc("/order", order)
	r.HandleFunc("/customers", customers)
	r.HandleFunc("/customers/{email}", customerByEmail)
	http.Handle("/", r)

	log.Fatal("listening: ", http.ListenAndServe(":7050", nil))
}
