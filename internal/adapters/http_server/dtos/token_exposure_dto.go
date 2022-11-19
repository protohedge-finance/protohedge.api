package http_server_dtos

type TokenExposureResponseDto struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Symbol string `json:"symbol"`
}
