package model

type UpdateOrderRequest struct {
	Schedule string `json:"schedule" binding:"required"`
}
