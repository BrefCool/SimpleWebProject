package routers

import (
	"myRepo/SimpleWebProject/ToDoList/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
