package server

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/services/cache"
	"github.com/YoungGoofy/WB_L0/internal/services/db"
	route "github.com/YoungGoofy/WB_L0/internal/services/server/handler"
	"github.com/YoungGoofy/WB_L0/internal/services/server/nats"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	config   *config.Config
	client   *pgxpool.Pool
	cache    *cache.Cache
	log      *log.Logger
	stanConn stan.Conn
}

func NewSubServer(config *config.Config, pool *pgxpool.Pool, cache *cache.Cache, log *log.Logger, conn stan.Conn) *Server {
	return &Server{config: config, client: pool, cache: cache, log: log, stanConn: conn}
}

func (s *Server) RunServer() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pgRepo := db.NewOrderRepository(s.client)
	cacheRepo := db.NewCache(s.cache)
	repo := db.NewRepositories(pgRepo, cacheRepo)

	var validate *validator.Validate
	validate = validator.New()

	go func() {
		orderSubscriber := nats.NewSubscriber(s.stanConn, s.log, validate, repo)
		orderSubscriber.Run(ctx)
	}()
	s.log.Println("Start working server")
	handler := route.NewHandler(*repo, s.log)
	handler.Route()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.log.Fatalf("ctx.Done: %v", done)
	}

	s.log.Println("Server Exited Property")

	return nil
}
