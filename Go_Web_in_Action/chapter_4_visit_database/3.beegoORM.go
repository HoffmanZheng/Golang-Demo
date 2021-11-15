package main

import (
	"fmt"

	"github.com/beego/beego/client/orm"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int
	Name       string
	Tel        string
	AvatarUrl  string
	Address    string
	CreatedAt  string
	ModifiedAt string
}

var o orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		"root:123456@tcp(127.0.0.1:3306)/golang-demo?charset=utf8")
	orm.RegisterModelWithPrefix("tb_", new(User)) // table name: tb_user
	fmt.Println("beego orm init finished!")
	o = orm.NewOrm()
}

func getUserById(id int) {
	fmt.Printf("get user by id, id: %d \n", id)
	user := &User{Id: id}
	err := o.Read(user)
	if err == orm.ErrNoRows {
		fmt.Println("no result found")
	} else if err == orm.ErrMissPK {
		fmt.Println("can't find primary key")
	} else {
		fmt.Printf("%v \n", *user)
	}
}

func insertNewUser() {
	u := new(User)
	u.Name = "Leo"
	u.Address = "Beijing"
	u.Tel = "155231455"
	u.CreatedAt = "2021-09-11 23:08"
	u.ModifiedAt = "2021-09-11 23:08"
	fmt.Println(o.Insert(u))
}

func updateUser() {
	user := new(User)
	user.Id = 1
	user.Name = "James"
	n, err := o.Update(user, "name")
	if err != nil {
		fmt.Printf("update failed, err: %v \n", err)
	} else {
		fmt.Printf("affect rows: %d \n", n)
	}
}

func main() {
	insertNewUser()
	getUserById(1)
	updateUser()
}
