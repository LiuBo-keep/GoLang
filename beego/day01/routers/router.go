package routers

import (
	"day01/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})

	web.Router("/find", &controllers.UserController{}, "get:FindAllUser")
}
