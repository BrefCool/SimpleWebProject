package models

import ( 
	"gopkg.in/mgo.v2/bson"

	"github.com/astaxie/beego"
)

// task model
type Task struct {
	Id		bson.ObjectId `bson:"_id" json:"id" form:"-"`
	Name	string		  `bson:"name" json:"name" form:"name"`
	Done	bool		  `bson:"done" json:"done" form:"done"`
}

// db & collection info
var (
	db = beego.AppConfig.String("db")
	collection = beego.AppConfig.String("collection")
)

func InsertTask(task Task) error {
	return Insert(db, collection, task)
} 

func FindAllTasks() ([]Task, error) {
	var result []Task
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func UpdateTask(task Task) error {
	return Update(db, collection, bson.M{"_id": task.Id}, task)
}

func RemoveTask(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}