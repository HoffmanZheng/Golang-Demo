package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // obj in this pkg not directly been used
)

type MyUser struct {
	Id         int
	Name       string
	Tel        string
	avatalUrl  string
	address    string
	createdAt  string
	modifiedAt string
}

var myDB *sql.DB
var user MyUser

func init() {
	// create DB instance
	myDB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golang-demo")
	err := myDB.Ping() // connect and verify the database connection
	if err != nil {
		fmt.Printf("error during db ping: %v \n", err)
	}
	fmt.Println("init DB successfully!")
}

func queryUserById(id int) {
	err := myDB.QueryRow("select id, name, tel, avatar_url, address, created_at, modified_at from `tb_user` where id=?",
		id).Scan(&user.Id, &user.Name, &user.Tel, &user.avatalUrl, &user.address, &user.createdAt, &user.modifiedAt)
	if err != nil {
		fmt.Printf("error during queryUserById, err: %v \n", err)
	}
	fmt.Printf("get user through queryUserById: %v \n", user)
}

func queryMultiRows() {
	rows, err := myDB.Query("select id, name, tel, avatar_url, address, created_at, modified_at from `tb_user` where id > ?", 0)
	if err != nil {
		fmt.Printf("error during queryMultiRows, err: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Tel, &user.avatalUrl, &user.address, &user.createdAt, &user.modifiedAt)
		if err != nil {
			fmt.Printf("error during queryMultiRows Scan, err: %v", err)
		}
		fmt.Printf("get user through queryMultiRows: %v \n", user)
	}
}

func insertRow() {
	res, err := myDB.Exec("insert into tb_user(name, tel, avatar_url, address, created_at, modified_at) values (?, ?, ?, ?, now(), now())",
		"Hoffman", "16621663666", "http://avatal2", "Shanghai")
	if err != nil {
		fmt.Printf("error during insertRow, err: %v", err)
	}
	id, err := res.LastInsertId() // get the last id which is just inserted
	if err != nil {
		fmt.Printf("error during getLasetInsertId, err: %v", err)
	}
	fmt.Printf("insert successfully, id: %d", id)
}

func updateRow() {
	res, err := myDB.Exec("update tb_user set tel = ? where id = ?", "1882188388", 1)
	if err != nil {
		fmt.Printf("error during updateRow, err: %v", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("error during updateRow, err: %v", err)
	}
	fmt.Printf("update successfully, affexted rows: %d \n", n)
}

func preparedQuery() {
	stmt, err := myDB.Prepare("select id, name, tel, avatar_url, address, created_at, modified_at from `tb_user` where id > ?")
	if err != nil {
		fmt.Printf("error during preparedQuery, err: %v", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("error during stmt.Query, err: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Tel, &user.avatalUrl, &user.address, &user.createdAt, &user.modifiedAt)
		if err != nil {
			fmt.Printf("error during preparedQuery Scan, err: %v", err)
		}
		fmt.Printf("get user through preparedQuery: %v \n", user)
	}
}

func main() {
	queryUserById(1)
	queryMultiRows()
	insertRow() // unique key produce duplicate entry exception
	updateRow()
	preparedQuery()
	defer myDB.Close()
}
