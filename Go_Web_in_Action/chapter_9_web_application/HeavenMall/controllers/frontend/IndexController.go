package frontend

import (
	"fmt"
	"time"

	"HeavenMall/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {

	//调用功能功能
	c.BaseInit()

	//开始时间
	startTime := time.Now().UnixNano()

	//获取轮播图 注意获取的时候要写地址
	banner := []models.Banner{}
	if hasBanner := models.CacheDb.Get(nil, "banner", &banner); hasBanner == true {
		c.Data["bannerList"] = banner
	} else {
		models.DB.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)
		c.Data["bannerList"] = banner
		models.CacheDb.Set(nil, "banner", banner)
	}

	//获取手机商品列表
	redisPhone := []models.Product{}
	if hasPhone := models.CacheDb.Get(nil, "phone", &redisPhone); hasPhone == true {
		c.Data["phoneList"] = redisPhone
	} else {
		phone := models.GetProductByCategory(1, "hot", 8)
		c.Data["phoneList"] = phone
		models.CacheDb.Set(nil, "phone", phone)
	}
	//获取电视商品列表
	redisTv := []models.Product{}
	if hasTv := models.CacheDb.Get(nil, "tv", &redisTv); hasTv == true {
		c.Data["tvList"] = redisTv
	} else {
		tv := models.GetProductByCategory(4, "best", 8)
		c.Data["tvList"] = tv
		models.CacheDb.Set(nil, "tv", tv)
	}

	//结束时间
	endTime := time.Now().UnixNano()

	fmt.Println("执行时间", endTime-startTime)

	c.TplName = "frontend/index/index.html"
}
