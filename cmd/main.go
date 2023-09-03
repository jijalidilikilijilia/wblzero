package main

import (
	"fmt"
	"log"

	"github.com/jijalidilikilijilia/wblzero/handlers"
	"github.com/jijalidilikilijilia/wblzero/internal/cache"
	"github.com/jijalidilikilijilia/wblzero/internal/database"
	"github.com/jijalidilikilijilia/wblzero/internal/nats"
	"github.com/jijalidilikilijilia/wblzero/web"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file in main: %s", err)
	}

	//NATS
	natsURL := viper.GetString("nats.url")
	natsClusterID := viper.GetString("nats.cluster_id")
	natsClientID := viper.GetString("nats.client_id")
	natsSubject := viper.GetString("nats.subject")
	natsQueueGroup := viper.GetString("nats.queueGroup")

	//DB
	dbDSN := viper.GetString("db.dsn")

	//REDIS
	cacheAddr := viper.GetString("cache.addr")
	cachePassword := viper.GetString("cache.password")
	cacheDB := viper.GetInt("cache.db")

	// Создание экземпляров
	dbHandler, err := database.NewDBHandler(dbDSN)
	if err != nil {
		log.Fatalf("Failed to create DB handler: %s", err)
	}
	log.Println("DB SUCCESS")
	defer dbHandler.Close()

	natsHandler, err := nats.NewNatsHandler(natsClusterID, natsClientID, natsURL)
	if err != nil {
		log.Fatalf("Failed to create NATS handler: %s", err)
	}
	log.Println("NATS SUCCESS")
	defer natsHandler.Close()

	cacheHandler, err := cache.NewCacheHandler(cacheAddr, cachePassword, cacheDB)
	if err != nil {
		log.Fatalf("Failed to create REDOS handler: %s", err)
	}
	log.Println("REDIS SUCCESS")
	defer cacheHandler.Close()

	//Замыкание
	orderHandler := handlers.CreateOrderSubscribeHandler(dbHandler, cacheHandler)
	sub, err := natsHandler.Subscribe(natsSubject, natsQueueGroup, orderHandler)
	if err != nil {
		log.Println("error sub ", err)
	}
	defer sub.Close()

	server := web.NewServer(cacheHandler, dbHandler, natsHandler)

	address := viper.GetString("server.address")
	fmt.Printf("Server is running at %s\n", address)
	if err := server.Run(address); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
