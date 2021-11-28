package routers

import (
	"HeavenMall/controllers"
	"HeavenMall/controllers/frontend"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &frontend.IndexController{})
}
