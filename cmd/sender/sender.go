package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/jijalidilikilijilia/wblzero/models"

	"github.com/nats-io/stan.go"
)

func main() {
	natsURL := "nats://localhost:4222"
	clusterID := "my_cluster"
	clientID := "my_client"

	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Printf("Failed to connect to NATS in SENDER: %s", err)
		return
	}
	defer conn.Close()

	var order models.Order
	data, err := os.ReadFile("testData.json")
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal([]byte(data), &order); err != nil {
		log.Printf("Failed to unmarshal JSON data: %s", err)
		return
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Printf("Failed to marshal order data in SENDER: %s", err)
		return
	}

	err = conn.Publish("wbl0", orderJSON)
	if err != nil {
		log.Printf("Failed to publish order in SENDER: %s", err)
		return
	}

	log.Println("Order published successfully")
	time.Sleep(time.Second) // Wait for a moment before exiting
}
