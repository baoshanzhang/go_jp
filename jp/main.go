// test
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

const (
	DBName   = "student"
	Host     = "localhost"
	User     = "root"
	Port     = 3306
	Password = ""
	Charset  = "utf8"
)

var engine *xorm.Engine

//database model
type Student struct {
	Id         string    `xorm:"id"`
	Name       string    `xorm:"name"`
	Age        int32     `xorm:"age"`
	Adress     string    `xorm:"adress"`
	CreateTime time.Time `xorm:"create_time"`
	UpdateTime time.Time `xorm:"update_time"`
	Version    int32     `xorm:"version"`
}

//插入数据到数据库
func InitDatabase() (err error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", User, Password, Host, Port, DBName, Charset)
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Failed to connect database", err)
		return
	}

	engines := make(map[string]*xorm.Engine, 0) //字典
	engines["rijin"] = engine
	for _, v := range engines {
		v.Logger().SetLevel(core.LOG_INFO)
	}
	return
}

func main() {
	err := InitDatabase()
	if err != nil {
		panic(err)
	}

	//插入单条
	student := &Student{
		Id:         "1234",
		Name:       "Frake",
		Age:        18,
		Adress:     "beijin",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if _, err = engine.InsertOne(student); err != nil {
		log.Fatal("Failed to insert student data to db", err)
		return
	}

}
