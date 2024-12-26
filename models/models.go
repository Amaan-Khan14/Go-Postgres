package models

type Stock struct {
	StockId int64  `json:"stock_id"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}
