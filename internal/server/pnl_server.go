package server

import (
	"encoding/json"
	"net/http"

	"github.com/gmx-delta-neutral/gmx-neutral.api/internal/pnl"
)

type pnlServer struct {
	pnlService pnl.Service
}

func NewPnlServer(pnlService pnl.Service) pnlServer {
	return pnlServer{pnlService: pnlService}
}

func (s *pnlServer) GetPnl(w http.ResponseWriter, r *http.Request) {
	profitAndLoss, err := s.pnlService.GetPnl()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(profitAndLoss)
}
