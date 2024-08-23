package model

type CreateOrderResponse struct {
	Id       int    `json:"id"`
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}
