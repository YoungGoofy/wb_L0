package main

import (
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/services"
	"github.com/YoungGoofy/WB_L0/internal/services/server"
	"log"
	"os"
)

func main() {
	var cfg config.Config
	_ = cfg.InitFromEnv()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("Start publish server")
	natsConn, err := services.NewNatsConnect(&cfg, "publisher", logger)
	if err != nil {
		logger.Println(err)
		return
	}
	ps := server.NewPubServer(logger, natsConn, &cfg)
	logger.Fatal(ps.Run())
}
