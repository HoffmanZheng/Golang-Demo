package models

import (
	_ "github.com/jinzhu/gorm"
)

type ProductImage struct {
	Id      int    `json:"id"`
	ProductId int    `json:"product_id"`
	ImgUrl  string `json:"img_url"`
	ColorId int    `json:"color_id"`
	Sort    int    `json:"sort"`
	AddTime int    `json:"add_time"`
	Status  int    `json:"status"`
}

func (ProductImage) TableName() string {
	return "product_image"
}
