package routers

import (
	"github.com/liuruichao555/beego-demo1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
