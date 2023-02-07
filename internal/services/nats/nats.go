package nats

import (
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

const (
	connectWait        = time.Second * 30
	pubAckWait         = time.Second * 30
	interval           = 10
	maxOut             = 5
	maxPubAcksInflight = 25
)

func NewNatsConnect(cfg *config.Config, clientID string, log *log.Logger) (stan.Conn, error) {
	return stan.Connect(
		cfg.NatsClusterId,
		clientID,
		stan.ConnectWait(connectWait),
		stan.PubAckWait(pubAckWait),
		stan.Pings(interval, maxOut),
		stan.MaxPubAcksInflight(maxPubAcksInflight),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.NatsURL("nats://"+cfg.NatsHostname+":4222"))
}
