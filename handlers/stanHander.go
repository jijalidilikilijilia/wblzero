package handlers

import (
	"encoding/json"
	"log"

	"github.com/jijalidilikilijilia/wblzero/internal/cache"
	"github.com/jijalidilikilijilia/wblzero/internal/database"
	"github.com/jijalidilikilijilia/wblzero/models"

	"github.com/nats-io/stan.go"
)

func CreateOrderSubscribeHandler(dbHandler *database.DBHandler, cacheHandler *cache.CacheHandler) stan.MsgHandler {
	return func(msg *stan.Msg) {
		log.Println("ПОЛУЧЕНО СООБЩЕНИЕ: ", string(msg.Data))
		var order models.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Fatal("Error unmarshaling order data. Error: ", err)
			return
		}

		if err := dbHandler.GetDB().Create(&order).Error; err != nil {
			log.Fatal("failed to save order to the database. Error: ", err)
			return
		}

		if err = cacheHandler.SetValue(order.OrderUID, string(msg.Data), 0); err != nil {
			log.Println("failed to set order to cache. Error: ", err)
			return
		}
	}
}
