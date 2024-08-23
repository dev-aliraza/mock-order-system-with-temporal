package model

type UpdateOrderResponse struct {
	Id       int    `json:"id"`
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}
