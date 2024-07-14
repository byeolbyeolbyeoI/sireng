package main

import (
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	"github.com/chaaaeeee/sireng/database"
	"github.com/chaaaeeee/sireng/server"
	"github.com/chaaaeeee/sireng/util"
)

func main() {
	conf := config.GetConfig()
	util := util.NewUtil(conf)

	db := database.NewDatabase(conf).GetDb()
	fmt.Println("Connected to the database...")
	// set error handling(?)
	server.NewServer(conf, db, util).Start()
}
