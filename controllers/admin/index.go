package admin

import (
	"github.com/astaxie/beego"
)

type AdminIndexController struct {
	beego.Controller
}

func (c *AdminIndexController) Get(){
	c.Data["Website"] = "dfdfdffdfd.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}