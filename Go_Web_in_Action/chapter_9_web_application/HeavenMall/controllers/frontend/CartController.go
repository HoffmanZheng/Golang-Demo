package frontend

import (
	"HeavenMall/models"
	"strconv"
)

//购物车结构体
type CartController struct {
	BaseController
}

//购物车展示
func (c *CartController) Get() {
	c.BaseInit()
	cartList := []models.Cart{}
	models.Cookie{}.Get(c.Ctx, "cartList", &cartList)

	var allPrice float64
	//执行计算总价
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	c.Data["cartList"] = cartList
	c.Data["allPrice"] = allPrice
	c.TplName = "frontend/cart/cart.html"
}

func (c *CartController) AddCart() {
	c.BaseInit()

	colorId, err1 := c.GetInt("color_id")
	productId, err2 := c.GetInt("product_id")

	product := models.Product{}
	productColor := models.ProductColor{}
	err3 := models.DB.Where("id=?", productId).Find(&product).Error
	err4 := models.DB.Where("id=?", colorId).Find(&productColor).Error

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {

		c.Ctx.Redirect(302, "/item_"+strconv.Itoa(product.Id)+".html")
		return
	}
	// 1.获取增加购物车的商品数据
	currentData := models.Cart{
		Id:             productId,
		Title:          product.Title,
		Price:          product.Price,
		ProductVersion: product.ProductVersion,
		Num:            1,
		ProductColor:   productColor.ColorName,
		ProductImg:     product.ProductImg,
		ProductGift:    product.ProductGift, //赠品
		ProductAttr:    "",                  //根据自己的需求拓展
		Checked:        true,                //默认选中
	}

	//2.判断购物车有没有数据（cookie）
	cartList := []models.Cart{}
	models.Cookie{}.Get(c.Ctx, "cartList", &cartList)
	if len(cartList) > 0 { //购物车有数据
		//4、判断购物车有没有当前数据
		if models.CartHasData(cartList, currentData) {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == currentData.Id && cartList[i].ProductColor == currentData.ProductColor && cartList[i].ProductAttr == currentData.ProductAttr {
					cartList[i].Num = cartList[i].Num + 1
				}
			}
		} else {
			cartList = append(cartList, currentData)
		}
		models.Cookie{}.Set(c.Ctx, "cartList", cartList)

	} else {
		//3.如果购物车没有任何数据，直接把当前数据写入cookie
		cartList = append(cartList, currentData)
		models.Cookie{}.Set(c.Ctx, "cartList", cartList)
	}

	c.Data["product"] = product
	c.TplName = "frontend/cart/addcart_success.html"
}

func (c *CartController) DecCart() {
	var flag bool
	var allPrice float64
	var currentAllPrice float64
	var num int

	productId, _ := c.GetInt("product_id")
	productColor := c.GetString("product_color")
	productAttr := ""

	cartList := []models.Cart{}
	models.Cookie{}.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			if cartList[i].Num > 1 {
				cartList[i].Num = cartList[i].Num - 1
			}
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		models.Cookie{}.Set(c.Ctx, "cartList", cartList)
		c.Data["json"] = map[string]interface{}{
			"success":         true,
			"message":         "修改数量成功",
			"allPrice":        allPrice,
			"currentAllPrice": currentAllPrice,
			"num":             num,
		}

	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
	}
	c.ServeJSON()
}

func (c *CartController) IncCart() {
	var flag bool
	var allPrice float64
	var currentAllPrice float64
	var num int

	productId, _ := c.GetInt("product_id")
	productColor := c.GetString("product_color")
	productAttr := ""

	cartList := []models.Cart{}
	models.Cookie{}.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			cartList[i].Num = cartList[i].Num + 1
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		models.Cookie{}.Set(c.Ctx, "cartList", cartList)
		c.Data["json"] = map[string]interface{}{
			"success":         true,
			"message":         "修改数量成功",
			"allPrice":        allPrice,
			"currentAllPrice": currentAllPrice,
			"num":             num,
		}

	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
	}
	c.ServeJSON()
}

func (c *CartController) ChangeOneCart() {
	var flag bool
	var allPrice float64
	productId, _ := c.GetInt("product_id")
	productColor := c.GetString("product_color")
	productAttr := ""

	cartList := []models.Cart{}
	models.Cookie{}.Get(c.Ctx, "cartList", &cartList)

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			cartList[i].Checked = !cartList[i].Checked
			flag = true
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		models.Cookie{}.Set(c.Ctx, "cartList", cartList)
		c.Data["json"] = map[string]interface{}{
			"success":  true,
			"message":  "修改状态成功",
			"allPrice": allPrice,
		}

	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
	}
	c.ServeJSON()
}

//全选反选
func (c *CartController) ChangeAllCart() {
	flag, _ := c.GetInt("flag")
	var allPrice float64
	cartList := []models.Cart{}
	models.Cookie{}.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if flag == 1 {
			cartList[i].Checked = true
		} else {
			cartList[i].Checked = false
		}
		//计算总价
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	models.Cookie{}.Set(c.Ctx, "cartList", cartList)

	c.Data["json"] = map[string]interface{}{
		"success":  true,
		"allPrice": allPrice,
	}
	c.ServeJSON()
}

func (c *CartController) DelCart() {
	productId, _ := c.GetInt("product_id")
	productColor := c.GetString("product_color")
	productAttr := ""

	cartList := []models.Cart{}
	models.Cookie{}.Get(c.Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			//执行删除
			cartList = append(cartList[:i], cartList[(i + 1):]...)
		}
	}
	models.Cookie{}.Set(c.Ctx, "cartList", cartList)

	c.Redirect("/cart", 302)
}
