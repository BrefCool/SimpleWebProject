package controllers

import (

	"encoding/json"

	"net/http"

	"ToDoList/models"
	
	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

type MainController struct {
	beego.Controller
}

// parse the task struct from request body
func (c *MainController) bind() (ta models.Task) {
	json.NewDecoder(c.Ctx.Request.Body).Decode(&ta)
	return
}

func (c *MainController) Home() {
	c.TplName = "index.html"
}

func (c *MainController) GetTasks() {
	datas, err := models.FindAllTasks()
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = datas
		c.ServeJSON()
	}
}

func (c *MainController) AddTask() {
	task := c.bind()
	task.Id = bson.NewObjectId()

	err := models.InsertTask(task)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = task
		c.ServeJSON()
	}
}

func (c *MainController) EditTask() {
	task := c.bind()

	err := models.UpdateTask(task)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = task
		c.ServeJSON()
	}
}

func (c *MainController) DeleteTask() {
	id := c.Ctx.Input.Param(":id")
	err := models.RemoveTask(id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = id
		c.ServeJSON()
	}
}
