package server

import (
	"encoding/json"
	"net/http"

	"github.com/protohedge/protohedge.api/internal/models/enums"
)

type ResponseType uint64

const (
	Success ResponseType = iota
	Error
)

func (s ResponseType) String() string {
	switch s {
	case Success:
		return "success"
	case Error:
		return "error"
	}
	return "unknown"
}

type Response struct {
	Type      string      `json:"type,omitempty"`
	ErrorCode string      `json:"errorCode,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func ErrorRes(w http.ResponseWriter) {
	ErrResWithCode(w, enums.ServerError)
}

func ErrResWithCode(w http.ResponseWriter, errorCode enums.Error) {
	w.WriteHeader(http.StatusInternalServerError)
	res := Response{Type: "error", ErrorCode: errorCode.String()}
	json.NewEncoder(w).Encode(res)
}

func SuccessResWithData(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	res := Response{
		Type: "success",
		Data: data,
	}
	json.NewEncoder(w).Encode(res)
}
