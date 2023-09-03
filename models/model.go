package models

import (
	"time"
)

// Order модель для заказа
type Order struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery" gorm:"foreignkey:OrderID"`
	Payment           Payment   `json:"payment" gorm:"foreignkey:OrderID"`
	Items             []Item    `json:"items" gorm:"foreignkey:OrderID"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey" gorm:"column:shardkey"`
	SMID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OOFShard          string    `json:"oof_shard"`
}

// Delivery модель для информации о доставке
type Delivery struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	OrderID uint   `json:"order_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

// Payment модель для информации о платеже
type Payment struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	OrderID      uint   `json:"order_id"`
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

// Item модель для информации о товаре
type Item struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	OrderID     uint   `json:"order_id"`
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid" gorm:"column:rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NMID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
