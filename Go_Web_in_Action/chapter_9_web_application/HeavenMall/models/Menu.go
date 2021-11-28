package models

import (
	_ "github.com/jinzhu/gorm"
)

type Menu struct {
	Id          int
	Title       string
	Link        string
	Position    int
	IsOpennew   int
	Relation    string
	Sort        int
	Status      int
	AddTime     int
	ProductItem []Product `gorm:"-"`
}

func (Menu) TableName() string {
	return "menu"
}
