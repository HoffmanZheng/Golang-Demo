package models

import (
	_ "github.com/jinzhu/gorm"
)

type ProductColor struct {
	Id         int
	ColorName  string
	ColorValue string
	Status     int
	Checked    bool `gorm:"-"`
}

func (ProductColor) TableName() string {
	return "product_color"
}
