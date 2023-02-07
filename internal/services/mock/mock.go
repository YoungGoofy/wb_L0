package mock

import (
	"github.com/YoungGoofy/WB_L0/internal/models"
	"github.com/brianvoe/gofakeit/v6"
)

func NewOrder() *models.Orders {
	var order models.Orders
	gofakeit.Seed(0)
	err := gofakeit.Struct(&order)
	if err != nil {
		return nil
	}
	return &order
}
