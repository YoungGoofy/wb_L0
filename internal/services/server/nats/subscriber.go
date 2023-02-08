package nats

import (
	"context"
	"encoding/json"
	"github.com/YoungGoofy/WB_L0/internal/models"
	"github.com/YoungGoofy/WB_L0/internal/services/db"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"log"
)

type Subscriber struct {
	log      *log.Logger
	mix      *db.Repositories
	stanConn stan.Conn
	validate *validator.Validate
}

func NewSubscriber(conn stan.Conn, logger *log.Logger, validate *validator.Validate, mix *db.Repositories) *Subscriber {
	return &Subscriber{log: logger, mix: mix, stanConn: conn, validate: validate}
}

func (s *Subscriber) createOrder(ctx context.Context) stan.MsgHandler {
	return func(msg *stan.Msg) {
		var order models.Orders

		err := s.validate.Struct(order)
		if err != nil {
			s.log.Printf("Validate error: %v", err)
			return
		}
		if err = json.Unmarshal(msg.Data, &order); err != nil {
			s.log.Printf("Unmarshal error: %v", err)
			return
		}
		if err = s.mix.Create(ctx, &order); err != nil {
			s.log.Printf("Create order error: %v", err)
			return
		}
	}
}

func (s *Subscriber) Run(ctx context.Context) {
	s.log.Println("Get data")
	go s.stanConn.Subscribe("order:create", s.createOrder(ctx))
}
