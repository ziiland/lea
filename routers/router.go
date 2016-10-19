package routers

import (
	"lea/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/land", &controllers.LoginController{})
    beego.Router(controllers.ZLD_PATH_LOGIN, &controllers.LoginController{})
    beego.Router(controllers.ZLD_PATH_WORKER, &controllers.WorkerController{})
    beego.Router(controllers.ZLD_PATH_TASK, &controllers.TaskController{})


}
