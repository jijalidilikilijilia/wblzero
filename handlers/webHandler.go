package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jijalidilikilijilia/wblzero/internal/cache"
	"github.com/jijalidilikilijilia/wblzero/internal/database"
	"github.com/jijalidilikilijilia/wblzero/models"

	"github.com/gin-gonic/gin"
)

func GetOrderHandler(db *database.DBHandler, cache *cache.CacheHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderUID := c.Param("uid")

		// Получаем данные из кэша
		cachedData, err := cache.GetValue(orderUID)
		if err == nil {
			log.Println("Данные из кеша")
			var order models.Order
			json.Unmarshal([]byte(cachedData), &order)
			c.HTML(http.StatusOK, "order.html", gin.H{
				"Order": order,
			})
			return
		}
		if err != nil {
			log.Printf("Данные о заказе %s не были найдены в кеше\n", orderUID)
		}
		// Получаем данные из БД и сохраняем в кеше
		var order models.Order
		result := db.GetDB().Table("wblzero_db.orders").Preload("Delivery").Preload("Payment").Preload("Items").Where("order_uid = ?", orderUID).First(&order)
		if result.Error != nil {
			c.String(http.StatusOK, "Заказ не найден")
			return
		}

		jsonOrderData, err := json.Marshal(order)
		if err != nil {
			log.Println("Failed to marshal orderData to JSON. Error: ", err)
			return
		}

		if err := cache.SetValue(orderUID, string(jsonOrderData), time.Minute); err != nil {
			log.Println("save db data to cache fail. Error: ", err)
		}

		log.Println("Данные из БД")
		c.HTML(http.StatusOK, "order.html", gin.H{
			"Order": order,
		})
	}
}
