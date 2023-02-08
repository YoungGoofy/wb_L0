package services

import (
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/nats-io/stan.go"
	"log"
)

const (
	interval = 10
	maxOut   = 5
)

func NewNatsConnect(cfg *config.Config, clientID string, log *log.Logger) (stan.Conn, error) {
	return stan.Connect(
		cfg.NatsClusterId,
		clientID,
		stan.ConnectWait(stan.DefaultConnectWait),
		stan.PubAckWait(stan.DefaultAckWait),
		stan.Pings(interval, maxOut),
		stan.MaxPubAcksInflight(stan.DefaultMaxPubAcksInflight),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.NatsURL(stan.DefaultNatsURL))
}
