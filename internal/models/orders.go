package models

import "time"

type Orders struct {
	OrderUID          string    `json:"order_uid," fake:"{regex:[0-9a-z]{19}}"`
	TrackNumber       string    `json:"track_number" fake:"{regex:[A-Z]{14}}"`
	Entry             string    `json:"entry" fake:"{regex:[A-Z]{4}}"`
	Locale            string    `json:"locale" fake:"{randomstring:[ru,en,it]}"`
	InternalSignature string    `json:"internal_signature" fake:"{lettern:8}"`
	CustomerId        string    `json:"customer_id" fake:"{lettern:4}"`
	DeliveryService   string    `json:"delivery_service" fake:"{lettern:5}"`
	Shardkey          string    `json:"shardkey" fake:"{regex:[0-9]{2}}"`
	SmId              int       `json:"sm_id" fake:"{number:1, 5000}"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard" fake:"{regex:[0-9]{2}}" `

	Delivery `json:"delivery" `
	Payment  `json:"payment" `
	Items    []Item `json:"items" fakesize:"2"`
}
