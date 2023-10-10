package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) FindAllUser() {
	u.Ctx.WriteString("hello world")
}

func (u *UserController) FindUserById() {
	u.Ctx.WriteString("GetUserById")
}
