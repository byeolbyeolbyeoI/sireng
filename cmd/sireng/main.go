package main

import (
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	"github.com/chaaaeeee/sireng/server"
)

func main() {
	conf := config.GetConfig()
	fmt.Println("Server running in :8080")
	server.NewServer(conf).Start()
}
