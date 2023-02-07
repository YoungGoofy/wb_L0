package server

import (
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/services/cache"
	"github.com/YoungGoofy/WB_L0/internal/services/db"
	route "github.com/YoungGoofy/WB_L0/internal/services/server/handler"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Server struct {
	config *config.Config
	client *pgxpool.Pool
	cache  *cache.Cache
	log    *log.Logger
}

func NewSubServer(config *config.Config, pool *pgxpool.Pool, cache *cache.Cache, log *log.Logger) *Server {
	return &Server{config: config, client: pool, cache: cache, log: log}
}

func (s *Server) RunServer() error {
	pgRepo := db.NewOrderRepository(s.client)
	cacheRepo := db.NewCache(s.cache)
	repo := db.NewRepositories(pgRepo, cacheRepo)

	//var validate *validator.Validate
	//validate = validator.New()
	handler := route.NewHandler(*repo, s.log)
	handler.Route()
	return nil
}
