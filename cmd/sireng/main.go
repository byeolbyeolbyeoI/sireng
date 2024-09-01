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
//	@description Sireng is a badass API service that is designed for people who's lost in life, can't do shit for themselves and having a hard time focusing. Sireng tracks your learning process and turns it into informative graphs for you to use it as a means of review. Sireng has a chat feature that you can use to ask for feedback from your friends or you can simply use it to wake your friend up and tell em to lock tf in.

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
