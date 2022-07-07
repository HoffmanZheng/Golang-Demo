package models

import (
	_ "github.com/jinzhu/gorm"
)

type OrderItem struct {
	Id           int
	OrderId      int
	Uid          int
	ProductTitle string
	ProductId    int
	ProductImg   string
	ProductPrice float64
	ProductNum   int
	ProductVersion string
	ProductColor   string
	AddTime      int
}

func (OrderItem) TableName() string {
	return "order_item"
}
