package main

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/postgresql"
	"github.com/YoungGoofy/WB_L0/internal/services/db"
	"github.com/YoungGoofy/WB_L0/internal/services/mock"
	"log"
)

func main() {
	var newCfg config.Config
	_ = newCfg.InitFromEnv()
	pool, err := postgresql.NewClient(context.Background(), &newCfg)
	if err != nil {
		log.Print(err)
		return
	}
	defer pool.Close()
	repo := db.NewRepository(pool)
	for i := 0; i < 5; i++ {
		newOrder := mock.NewOrder()
		_, err = repo.CreateOrder(context.Background(), newOrder)
		if err != nil {
			log.Println(err)
			return
		}
	}
	//var name string
	//err = pool.QueryRow(context.Background(), testQuery, 0).Scan(&name)
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	//fmt.Println(name)
}
