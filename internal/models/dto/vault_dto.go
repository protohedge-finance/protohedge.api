package dto

type VaultResponseDto struct {
	VaultAddress       string                       `json:"vaultAddress"`
	PositionManagers   []PositionManagerResponseDto `json:"positionManagers"`
	VaultWorth         string                       `json:"vaultWorth"`
	AvailableLiquidity string                       `json:"availableLiquidity"`
}
