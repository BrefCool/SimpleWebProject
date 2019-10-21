package models

import (
	"log"

	"gopkg.in/mgo.v2"

	"github.com/astaxie/beego"
)

var (
	host = beego.AppConfig.String("host")
	user = beego.AppConfig.String("user")
	pass = beego.AppConfig.String("pass")
)

var globalS *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Username: user,
		Password: pass,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	//s, err := mgo.Dial(host)
	if err != nil {
		log.Fatalln("create session error ", err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}