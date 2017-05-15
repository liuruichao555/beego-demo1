package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "hello"
	c.Data["Email"] = "499244831@gmail.com"
	c.TplName = "index.tpl"
}