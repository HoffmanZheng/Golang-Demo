package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MyUser struct {
	Id         int       `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Tel        string    `gorm:"column:tel"`
	AvatarUrl  string    `gorm:"column:avatar_url"`
	Address    string    `gorm:"column:address"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	ModifiedAt time.Time `gorm:"column:modified_at"`
}

var (
	db            *gorm.DB
	sqlConnection = "root:123456@tcp(127.0.0.1:3306)/wxshop?charset=utf8&parseTime=true"
)

func init() {
	var err error
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		fmt.Printf("error during open sqlConnection, err: %v \n", err)
	}
	db.AutoMigrate(&MyUser{})
}

func main() {
	engine := gin.Default()
	userV2 := engine.Group("/api/v2/user")
	userV2.GET("/", getAllUser)
	userV2.GET("/:id", getUserById)
	userV2.POST("/", createUser)
	userV2.DELETE("/:id", deleteUserById)
	userV2.PUT("/:id", updateUserById)
	engine.Run()
}

func (u *MyUser) TableName() string {
	return "tb_user"
}

func getAllUser(c *gin.Context) {
	var users []MyUser
	db.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No user found",
		})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   users,
		})
}

func getUserById(c *gin.Context) {
	var user MyUser
	id := c.Param("id")
	db.First(&user, id)

	if user.Id == 0 { // the user does not exist
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No user found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   user,
		})
	}
}

func createUser(c *gin.Context) {
	name := c.PostForm("name")
	tel := c.PostForm("tel")
	avatalUrl := c.PostForm("avatalUrl")
	address := c.PostForm("address")
	fmt.Printf("name: %s, tel: %s. avatalUrl: %s, address: %s", name, tel, avatalUrl, address)
	user := MyUser{Name: name, Tel: tel, AvatarUrl: avatalUrl, Address: address, CreatedAt: time.Now(), ModifiedAt: time.Now()}
	db.Save(&user)
	c.JSON(201, user)
}

func deleteUserById(c *gin.Context) {
	var user MyUser
	id := c.Param("id")
	db.First(&user, id)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No user found!",
		})
		return
	} else {
		db.Delete(&user)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User deleted successfully!",
		})
	}
}

func updateUserById(c *gin.Context) {
	var user MyUser
	id := c.Param("id")
	db.First(&user, id)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No user found!",
		})
		return
	} else {
		db.Model(&user).Update("tel", c.PostForm("tel"))
		db.Model(&user).Update("name", c.PostForm("name"))
		db.Model(&user).Update("avatarUrl", c.PostForm("avatarUrl"))
		db.Model(&user).Update("address", c.PostForm("address"))
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Updated User successfully!",
		})
	}
}
