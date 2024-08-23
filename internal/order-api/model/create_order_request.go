package model

type CreateOrderRequest struct {
	Schedule string `json:"schedule" binding:"required"`
}
