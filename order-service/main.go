package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

type Order struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

var natsConn *nats.Conn

func main() {
	setupNATSProducer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order Service"))
	})
	http.HandleFunc("/placeOrder", placeOrderHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func setupNATSProducer() {
	var err error
	natsConn, err = nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal("Error connecting to NATS:", err)
	}
	log.Println("Connected to NATS successfully")
}

func placeOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	publishOrderToNATS(order)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order placed successfully"))
}

func publishOrderToNATS(order Order) {
	message, err := json.Marshal(order)
	if err != nil {
		log.Println("Error marshaling order:", err)
		return
	}

	err = natsConn.Publish("orders", message)
	if err != nil {
		log.Println("Error publishing message to NATS:", err)
		return
	}
	log.Println("Order published to NATS successfully")
}
