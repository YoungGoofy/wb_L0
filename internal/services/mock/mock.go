package mock

import (
	"github.com/YoungGoofy/WB_L0/internal/models"
	"github.com/brianvoe/gofakeit/v6"
)

func NewOrder() *models.Orders {
	var order models.Orders
	gofakeit.Seed(1)
	gofakeit.Struct(&order)
	return &order
}
