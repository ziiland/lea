package controllers

import (
	"github.com/astaxie/beego"
)

const ZLD_LOGIN_PATH string  = "/land/login"
const ZLD_WORKER_PATH string = "/land/worker"
const ZLD_TASK_PATH string = "/land/task"

const ZLD_PARA_WORKERID string = "WorkerId"
const ZLD_PARA_PWD string = "Password"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
