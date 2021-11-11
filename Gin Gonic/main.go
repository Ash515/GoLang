package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User      //nil
var Users2 = []User{} //[]

func main() {
	r := gin.Default()

	r.GET("/", getHandler)
	r.GET("/guser", GetUser)
	r.POST("/insert", CreateUser)
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}

}

func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":   "Ashwin",
		"Gender": "Male",
	})
}

func GetUser(d *gin.Context) {
	d.JSON(200, Users)
}

func CreateUser(e *gin.Context) {
	var reqbody User
	if err := e.ShouldBind(reqbody); err != nil {
		e.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request message",
		})
		return
	}
	Users = append(Users, reqbody)
	e.JSON(200, gin.H{
		"error":   false,
		"message": "Success!",
	})

}
