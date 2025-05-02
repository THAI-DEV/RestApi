package model

import "time"

type BuyProductRequestBody struct {
	ShoppingDate string `json:"shoppingDate"`
}

type ProductType struct {
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	NumUnit          int        `json:"numUnit"`
	UnitName         string     `json:"unitName"`
	MinimumUnit      int        `json:"minimumUnit"`
	RecordDate       *time.Time `json:"recordDate,omitempty"`
	ExpireDate       *time.Time `json:"expireDate,omitempty"`
	MinimumExpireDay int        `json:"minimumExpireDay"`
	BuyFlag          bool       `json:"buyFlag"`
}

type ProductListType []ProductType

type ResultBuyProduct struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	IsBuy       bool   `json:"isBuy"`
	Reason      string `json:"reason,omitempty"`
}

type ResultListBuyProduct []ResultBuyProduct
