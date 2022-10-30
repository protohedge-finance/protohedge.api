package main

import (
	"github.com/protohedge/protohedge.api/internal/config"
	"github.com/protohedge/protohedge.api/internal/server"
)

func main() {
	config := config.NewConfig()
	server.CreateServer(config)
}
