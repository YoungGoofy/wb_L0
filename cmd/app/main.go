package main

import (
	"context"
	"fmt"
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/postgresql"

	"log"
)

const (
	testQuery = "select name from parents where id=$1"
)

func main() {
	newCfg := config.Config{}
	err := newCfg.InitFromEnv()
	if err != nil {
		log.Print(err)
		return
	}
	pool, err := postgresql.NewClient(context.Background(), &newCfg)
	if err != nil {
		log.Print(err)
		return
	}
	defer pool.Close()
	var name string
	err = pool.QueryRow(context.Background(), testQuery, 0).Scan(&name)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println(name)
}
