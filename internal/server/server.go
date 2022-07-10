package server

import (
	"log"
	"net/http"

	"github.com/gmx-delta-neutral/gmx-neutral.api/internal/client"
	"github.com/gmx-delta-neutral/gmx-neutral.api/internal/pnl"
)

var port = ":8080"

func CreateServer() {
	queryClient, _, err := client.NewQueryClient()

	if err != nil {
		log.Fatal(err)
	}

	pnlService := pnl.NewPnlService(queryClient)
	pnlServer := NewPnlServer(pnlService)

	http.HandleFunc("/pnl", pnlServer.GetPnl)
	log.Printf("Listening on port %s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
