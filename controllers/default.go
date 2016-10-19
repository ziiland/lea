package controllers

import (
	"github.com/astaxie/beego"
)

const ZLD_PATH_LOGIN string  = "/land/login"
const ZLD_PATH_WORKER string = "/land/worker"
const ZLD_PATH_TASK string = "/land/task"

const ZLD_CMD_LOAD_PARA string = "LoadPara"
const ZLD_CMD_LOAD_VER string = "LoadVer"
const ZLD_CMD_UNLOAD string = "unload"

const ZLD_PARA_COMMAND string = "Command"
const ZLD_PARA_WORKERID string = "WorkerId"
const ZLD_PARA_PWD string = "Password"

const ZLD_STR_ON string = "on"
const ZLD_STR_OFF string = "off"
const ZLD_STR_OK string = "ok"

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
