package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/core/topology"
)

type customer struct {
	Email                   string `bson:"_id"` // using email as primary key
	RewardPoints            float64
	RewardsTier             string // TODO make this an enum for type safety and data usage
	RewardsTierName         string
	NextRewardsTier         string
	NextRewardsTierName     string
	NextRewardsTierProgress float64
}

func customers(w http.ResponseWriter, r *http.Request) {

	log.Println("getting customers")

	cur, err := collection.Find(context.TODO(), nil)
	if err != nil {
		log.Println("finding customers:", err)
		writeHTTPErr(w, err)
		return
	}
	defer cur.Close(context.TODO())

	var custs []customer
	for cur.Next(context.TODO()) {
		var cust customer
		err := cur.Decode(&cust)
		if err != nil {
			log.Println("decoding customer:", err)
			return
		}
		custs = append(custs, cust)
	}

	log.Printf("customers: %+v", custs)
	json.NewEncoder(w).Encode(custs)
}

func customerByEmail(w http.ResponseWriter, r *http.Request) {

	log.Println("getting customer")

	cust, err := getCustomer(mux.Vars(r)["email"])
	if err != nil {
		log.Println("getting customer:", err)
		writeHTTPErr(w, err)
		return
	}

	log.Printf("customer: %+v", cust)
	json.NewEncoder(w).Encode(cust)
}

func writeHTTPErr(w http.ResponseWriter, err error) {

	log.Println("writing http status for err:", err)

	if err == topology.ErrServerSelectionTimeout {
		w.WriteHeader(http.StatusGatewayTimeout)
	} else if strings.Contains(err.Error(), "mail:") {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getCustomer(email string) (*customer, error) {
	var cust customer
	res := collection.FindOne(context.TODO(), bson.NewDocument(bson.EC.String("_id", email)))
	if err := res.Decode(&cust); err != nil {
		log.Println("decoding customer:", err)
		return nil, err
	}
	return &cust, nil
}

func createCustomer(tx transaction) error {

	email := tx.Email

	// TODO check for valid email
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("creating customer:", err)
		return err
	}

	cust := customer{
		Email: email,
	}
	updateCustomer(&cust, tx.Total)

	res, err := collection.InsertOne(context.TODO(), cust)
	if err != nil {
		log.Println("creating customer:", err)
		return err
	}
	log.Printf("insert result: %+v\n", res)
	return nil
}

func updateCustomer(cust *customer, orderTotal float64) {

	cust.RewardPoints += orderTotal

	if cust.RewardPoints < 100 {
		cust.RewardsTier = ""
		cust.RewardsTierName = ""
		cust.NextRewardsTier = "A"
		cust.NextRewardsTierName = "5% off purchase"
		cust.NextRewardsTierProgress = cust.RewardPoints / 100
	} else if cust.RewardPoints < 200 {
		cust.RewardsTier = "A"
		cust.RewardsTierName = "5% off purchase"
		cust.NextRewardsTier = "B"
		cust.NextRewardsTierName = "10% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 100) / 100
	} else if cust.RewardPoints < 300 {
		cust.RewardsTier = "B"
		cust.RewardsTierName = "10% off purchase"
		cust.NextRewardsTier = "C"
		cust.NextRewardsTierName = "15% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 200) / 100
	} else if cust.RewardPoints < 400 {
		cust.RewardsTier = "C"
		cust.RewardsTierName = "15% off purchase"
		cust.NextRewardsTier = "D"
		cust.NextRewardsTierName = "20% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 300) / 100
	} else if cust.RewardPoints < 500 {
		cust.RewardsTier = "D"
		cust.RewardsTierName = "20% off purchase"
		cust.NextRewardsTier = "E"
		cust.NextRewardsTierName = "25% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 400) / 100
	} else if cust.RewardPoints < 600 {
		cust.RewardsTier = "E"
		cust.RewardsTierName = "25% off purchase"
		cust.NextRewardsTier = "F"
		cust.NextRewardsTierName = "30% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 500) / 100
	} else if cust.RewardPoints < 700 {
		cust.RewardsTier = "F"
		cust.RewardsTierName = "30% off purchase"
		cust.NextRewardsTier = "G"
		cust.NextRewardsTierName = "35% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 600) / 100
	} else if cust.RewardPoints < 800 {
		cust.RewardsTier = "G"
		cust.RewardsTierName = "35% off purchase"
		cust.NextRewardsTier = "H"
		cust.NextRewardsTierName = "40% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 700) / 100
	} else if cust.RewardPoints < 900 {
		cust.RewardsTier = "H"
		cust.RewardsTierName = "40% off purchase"
		cust.NextRewardsTier = "I"
		cust.NextRewardsTierName = "45% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 800) / 100
	} else if cust.RewardPoints < 1000 {
		cust.RewardsTier = "I"
		cust.RewardsTierName = "45% off purchase"
		cust.NextRewardsTier = "J"
		cust.NextRewardsTierName = "50% off purchase"
		cust.NextRewardsTierProgress = (cust.RewardPoints - 900) / 100
	} else {
		cust.RewardsTier = "J"
		cust.RewardsTierName = "50% off purchase"
		cust.NextRewardsTier = ""
		cust.NextRewardsTierName = ""
		cust.NextRewardsTierProgress = 0
	}
}
