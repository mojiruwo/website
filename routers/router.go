package routers

import (
	"website/controllers"
	"website/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &admin.AdminIndexController{})
}
