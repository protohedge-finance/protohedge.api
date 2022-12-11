package http_server_dtos

import "time"

type TimePointDto struct {
	Timestamp time.Time `json:"timestamp"`
	Point     string    `json:"point"`
}
