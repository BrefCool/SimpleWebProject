package routers

import (
	"ToDoList/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "Get:Home")
	beego.Router("/tasks/list", &controllers.MainController{}, "Get:GetTasks")
	beego.Router("/tasks/add", &controllers.MainController{}, "Post:AddTask")
	beego.Router("/tasks/edit", &controllers.MainController{}, "Put:EditTask")
	beego.Router("/tasks/:id", &controllers.MainController{}, "Delete:DeleteTask")
}
