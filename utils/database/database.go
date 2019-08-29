package database

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var(
	instance *gorm.DB
	once sync.Once
)

func Connent() *gorm.DB {
	db, err := gorm.Open("mysql","root:root123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}
func GetInstance() *gorm.DB {
	once.Do(func() {
		instance = Connent()
		instance.LogMode(true)
	})
	return instance
}
func Close()  {
	if instance == nil {
		return
	}
	instance.Close()
}
