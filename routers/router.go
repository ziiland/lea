package routers

import (
	"lea/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/land", &controllers.LoginController{})
    beego.Router(controllers.ZLD_LOGIN_PATH, &controllers.LoginController{})
    beego.Router(controllers.ZLD_WORKER_PATH, &controllers.WorkerController{})
    beego.Router(controllers.ZLD_TASK_PATH, &controllers.TaskController{})


}
