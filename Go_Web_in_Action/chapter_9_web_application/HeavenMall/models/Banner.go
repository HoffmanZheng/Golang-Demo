package models

import (
	_ "github.com/jinzhu/gorm"
)

type Banner struct {
	Id         int
	Title      string
	BannerType int
	BannerImg  string
	Link       string
	Sort       int
	Status     int
	AddTime    int
}

func (Banner) TableName() string {
	return "banner"
}
