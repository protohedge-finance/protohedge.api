package main

import (
	"github.com/protohedge/protohedge.api/internal/adapters/http_server"
	"github.com/protohedge/protohedge.api/internal/config"
)

func main() {
	config := config.NewConfig()

	http_server.CreateServer(config)
}
