package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRespon struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	println(password)
	c.JSON(http.StatusOK, UserRespon{
		StatusCode: 0,
		StatusMsg:  "",
		UserID:     1,
		UserName:   username,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	println(password)
	c.JSON(http.StatusOK, UserRespon{
		StatusCode: 0,
		StatusMsg:  "",
		UserID:     1,
		UserName:   username,
	})
}
