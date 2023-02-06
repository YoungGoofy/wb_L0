package main

import (
	"context"
	"fmt"
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/services/db"
	"github.com/YoungGoofy/WB_L0/internal/services/postgresql"
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
	repo := db.NewRepository(pool, pool)
	//for i := 0; i < 5; i++ {
	//	newOrder := mock.NewOrder()
	//	err = repo.Create(context.Background(), newOrder)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//}
	order, err := repo.GetOrderById(context.Background(), "1")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(order)
}
