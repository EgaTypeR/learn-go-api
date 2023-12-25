package model

type Product struct {
	ProductId   int    `json:"productId"`
	BarCode     string `json:"carCode"`
	ProductName string `json:"productName"`
	CategoryId  int    `json:"categoryId"`
	Price       int    `json:"price"`
}