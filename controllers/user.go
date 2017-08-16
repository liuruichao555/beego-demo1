package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Username string
	Age      int
}

func (c *UserController) Get() {
	user := User{Username: "刘瑞超", Age: 20}
	jsonStr, err := json.Marshal(user)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = string(jsonStr)
	}
	c.ServeJSON(true)
}