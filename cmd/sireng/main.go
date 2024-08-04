package main

import (
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	"github.com/chaaaeeee/sireng/database"
	"github.com/chaaaeeee/sireng/server"
	"github.com/chaaaeeee/sireng/util"
)

//	@title			Sireng API Documentation
//	@version		0.7
//	@description	Sinau bareng, a program that tracks user's study hours and share it with their friends! with Sireng, learning thins will be a lot more fun!

//	@contact.name	chaaaeeee
//	@contact.email	mash1o1o1o1@gmail.com

//	@host		localhost:8080

func main() {
	conf := config.GetConfig()
	util := util.NewUtil(conf)

	db := database.NewDatabase(conf).GetDb()
	fmt.Println("Connected to the database...")
	// set error handling(?)
	server.NewServer(conf, db, util).Start()
}
