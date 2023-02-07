package main

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/config"
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
	//cacheRepo := db.NewCache(&c)
	//pgRepo := db.NewOrderRepository(pool)
	//repo := db.NewRepositories(pgRepo, cacheRepo)
	//order, err := repo.GetById(ctx, "bn0joxy0phfkx9e75h9")
	//if err != nil {
	//	return
	//}
	//fmt.Println(order)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	subServer := server.NewSubServer(&cfg, pool, &c, logger)
	log.Println("http://localhost:8080/api")
	err = subServer.RunServer()
	if err != nil {
		log.Print(err)
		return
	}

}
