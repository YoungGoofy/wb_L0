package models

type Payment struct {
	Transaction  string `json:"transaction" fake:""`
	RequestId    string `json:"request_id" fake:""`
	Currency     string `json:"currency" fake:""`
	Provider     string `json:"provider" fake:""`
	Amount       int    `json:"amount" fake:"{number:1, 2389212}"`
	PaymentDt    int    `json:"payment_dt" fake:"{number:1, 2389212}"`
	Bank         string `json:"bank" fake:""`
	DeliveryCost int    `json:"delivery_cost" fake:"{number:1, 2389212}"`
	GoodsTotal   int    `json:"goods_total" fake:"{number:1, 2389212}"`
	CustomFee    int    `json:"custom_fee" fake:"{number:1, 2389212}"`
}
