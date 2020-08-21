package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type People struct{
	Name string `gorm:"name""`
	Age int `gorm:"age"`
	Sex string `gorm:"sex"`
}

func main() {
	db, err := gorm.Open("mysql", "root:@/gorm_example?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("fail to connect db, err:", err)
	}
	defer db.Close()
	db.Create(&People{"Jane",28,"F"})
}
