package server

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/services/server/nats"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type PubServer struct {
	log      *log.Logger
	natsConn stan.Conn
	cfg      *config.Config
}

func NewPubServer(logger *log.Logger, conn stan.Conn, cfg *config.Config) *PubServer {
	return &PubServer{log: logger, natsConn: conn, cfg: cfg}
}

func (ps *PubServer) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		orderPublisher := nats.NewPublisher(ps.natsConn, ps.log)
		orderPublisher.PublishOrder()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}

	ps.log.Println("Server Exited Property")

	return nil
}
