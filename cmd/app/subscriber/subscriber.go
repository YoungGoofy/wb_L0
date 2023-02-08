package main

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/services"
	"github.com/YoungGoofy/WB_L0/internal/services/cache"
	"github.com/YoungGoofy/WB_L0/internal/services/postgresql"
	"github.com/YoungGoofy/WB_L0/internal/services/server"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var cfg config.Config
	_ = cfg.InitFromEnv()
	pool, err := postgresql.NewClient(ctx, &cfg)
	if err != nil {
		log.Print(err)
		return
	}
	var c cache.Cache
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("Start subscriber")
	natsConn, err := services.NewNatsConnect(&cfg, "subscriber", logger)
	if err != nil {
		log.Println(err)
		return
	}
	subServer := server.NewSubServer(&cfg, pool, &c, logger, natsConn)
	log.Println("http://localhost:8080/api")
	err = subServer.RunServer()
	if err != nil {
		log.Print(err)
		return
	}
}
