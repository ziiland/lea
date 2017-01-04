package routers

import (
	"lea/common"
	"lea/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/land", &controllers.LoginController{})
    beego.Router(common.ZLD_PATH_LOGIN, &controllers.LoginController{})
    beego.Router(common.ZLD_PATH_WORKER, &controllers.WorkerController{})
    beego.Router(common.ZLD_PATH_TASK, &controllers.TaskController{})
    beego.Router(common.ZLD_PATH_PRICE, &controllers.PriceController{})
    beego.Router(common.ZLD_PATH_CELL, &controllers.CellController{})
    beego.Router(common.ZLD_PATH_PACKET, &controllers.PacketController{})
}
