package models

import (
	_ "github.com/jinzhu/gorm"
)

type Product struct {
	Id              int
	Title           string
	SubTitle        string
	ProductSn       string
	CateId          int
	ClickCount      int
	ProductNumber   int
	Price           float64
	MarketPrice     float64
	RelationProduct string
	ProductAttr     string
	ProductVersion  string
	ProductImg      string
	ProductGift     string
	ProductFitting  string
	ProductColor    string
	ProductKeywords string
	ProductDesc     string
	ProductContent  string
	IsDelete        int
	IsHot           int
	IsBest          int
	IsNew           int
	ProductTypeId   int
	Sort            int
	Status          int
	AddTime         int
}

func (Product) TableName() string {
	return "product"
}

func GetProductByCategory(cateId int, productType string, limitNum int) []Product {
	product := []Product{}
	productCate := []ProductCate{}
	DB.Where("pid=?", cateId).Find(&productCate)
	var tempSlice []int
	if len(productCate) > 0 {
		for i := 0; i < len(productCate); i++ {
			tempSlice = append(tempSlice, productCate[i].Id)
		}
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	switch productType {
	case "hot":
		where += "AND is_hot=1"
	case "best":
		where += "AND is_best=1"
	case "new":
		where += "AND is_new=1"
	default:
		break
	}
	DB.Where(where, tempSlice).Select("id,title,price,product_img,sub_title").Limit(limitNum).Order("sort desc").Find(&product)
	return product
}
