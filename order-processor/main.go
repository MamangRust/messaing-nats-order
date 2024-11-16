package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Order struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func main() {
	setupNATSConsumer()
}

func setupNATSConsumer() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal("Error connecting to NATS:", err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("orders", func(msg *nats.Msg) {
		var order Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println("Error unmarshaling order:", err)
			return
		}
		processOrder(order)
	})
	if err != nil {
		log.Fatal("Error subscribing to orders:", err)
	}

	log.Println("NATS consumer is now consuming messages from the 'orders' topic")
	select {}
}

func processOrder(order Order) {
	fmt.Printf("Processing Order ID: %d, Status: %s\n", order.ID, order.Status)

	time.Sleep(2 * time.Second) // Simulate order processing

	order.Status = "processed"
	log.Printf("Order ID %d processed successfully\n", order.ID)
}
