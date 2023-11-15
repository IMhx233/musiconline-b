package database

import (
	"fmt"
	"musiconline/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Login() {
	var UserLoginInfo model.UserInfo
	// UserLoginInfo.UserId = 1000
	 UserLoginInfo.Username = "abctffff"
	// UserLoginInfo.Password = "123456789"
	db, err := gorm.Open(mysql.Open("user:ubuntu@tcp(127.0.0.1:3306)/MusicOnline?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	a := db.First(&UserLoginInfo)
	fmt.Println(a.RowsAffected)

}
