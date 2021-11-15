package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type GormUser struct {
	ID       uint   `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", `root:123456@(127.0.0.1:3306)/golang-demo`)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// setup log
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	db.AutoMigrate(&GormUser{}) // create table

	GormUser := GormUser{
		Phone:    "13888888888",
		Name:     "hoffman",
		Password: md5Password("666666"),
	}

	insertUser(&GormUser)

	getUserByPhone(&GormUser, "13888888888")

	updatePhoneByUser(&GormUser, "18888888888")

	deleteUserByPhone(&GormUser, "18888888888")

	transaction(&GormUser, "13888888888")

}

func deleteUser(db *gorm.DB, gormUser *GormUser, s string) {
	panic("unimplemented")
}

func md5Password(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func insertUser(user *GormUser) {
	db.Save(user)
}

func deleteUserByPhone(user *GormUser, phone string) {
	db.Where("phone = ?", phone).Delete(user)
}

func getUserByPhone(user *GormUser, phone string) {
	db.Where("phone = ?", phone).Find(user)
	//db.First(&GormUser, "phone = ?", "18888888888")
	fmt.Println(user)
}

func updatePhoneByUser(user *GormUser, phone string) {
	err := db.Model(user).Where("phone = ?",
		"18888888888").Update("phone", "13888888888").Error
	if err != nil {
		panic(err)
	}
}

func transaction(user *GormUser, phone string) {
	tx := db.Begin() // start a transaction
	/*
		GormUser := GormUser{
			Phone:    "18888888888",
			Name:     "Shirdon",
			Password: md5Password("666666"),
		}
	*/
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	db.First(user, "phone = ?", phone)
	tx.Commit()
}
