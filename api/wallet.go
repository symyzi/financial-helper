package api

type createWalletRequest struct {
	string   `json:"wallet" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=RUB USD EUR"`
}
