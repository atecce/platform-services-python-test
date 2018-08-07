package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type transaction struct {
	Email string
	Total float64
}

func order(w http.ResponseWriter, r *http.Request) {

	log.Println("making order")

	var tx transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		log.Println("decoding order:", err)
	}
	log.Printf("tx: %+v\n", tx)

	cust, err := getCustomer(tx.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("customer does not exist. creating")
			createCustomer(tx)
		} else {
			log.Println("getting customer:", err)
			writeHTTPErr(w, err)
		}
		return
	}
	log.Printf("cust: %+v\n", cust)

	updateCustomer(cust, tx.Total)
	res := collection.FindOneAndReplace(context.TODO(), bson.NewDocument(bson.EC.String("_id", tx.Email)), cust)
	log.Printf("res: %+v\n", res)
}
