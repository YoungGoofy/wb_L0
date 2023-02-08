package nats

import (
	"encoding/json"
	"github.com/YoungGoofy/WB_L0/internal/services/mock"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

type Publisher struct {
	stanConn stan.Conn
	log      *log.Logger
}

func NewPublisher(conn stan.Conn, logger *log.Logger) *Publisher {
	return &Publisher{stanConn: conn, log: logger}
}

func (p *Publisher) PublishOrder() {
	for {
		randOrder := mock.NewOrder()
		byteOrder, err := json.Marshal(randOrder)
		if err != nil {
			p.log.Println(err)
		}

		err = p.stanConn.Publish("order:create", byteOrder)
		if err != nil {
			p.log.Println(err)
		}
		p.log.Println("Publish random order")
		time.Sleep(5 * time.Second)
	}
}
