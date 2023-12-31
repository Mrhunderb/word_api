package handler

import (
	"net/http"
	"strconv"

	"github.com/actionX/api/db"
	"github.com/gin-gonic/gin"
)

type UserRespon struct {
	StatusCode int
	StatusMsg  string
	UserID     int
	UserName   string
}

type InfoRespon struct {
	StatusCode int
	StatusMsg  string
	UserID     int
	UserName   string
	PlanID     int
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	new_user, err := db.InsertUser(username, password)
	if new_user == nil {
		c.JSON(http.StatusOK, UserRespon{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			UserID:     0,
			UserName:   "",
		})
		return
	}
	c.JSON(http.StatusOK, UserRespon{
		StatusCode: 0,
		StatusMsg:  "",
		UserID:     int(new_user.UserID),
		UserName:   new_user.UserName,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := db.FindUser(username, password)
	if user == nil {
		c.JSON(http.StatusOK, UserRespon{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			UserID:     0,
			UserName:   "",
		})
		return
	}
	c.JSON(http.StatusOK, UserRespon{
		StatusCode: 0,
		StatusMsg:  "",
		UserID:     int(user.UserID),
		UserName:   user.UserName,
	})
}

func GetInfo(c *gin.Context) {
	user_id := c.Query("user_id")
	user, err := db.FindUserByID(user_id)
	if user == nil {
		c.JSON(http.StatusOK, InfoRespon{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			UserID:     0,
			UserName:   "",
			PlanID:     0,
		})
		return
	}
	c.JSON(http.StatusOK, InfoRespon{
		StatusCode: 0,
		StatusMsg:  "",
		UserID:     int(user.UserID),
		UserName:   user.UserName,
		PlanID:     int(user.PlanID),
	})
}

func GetPlan(c *gin.Context) {
	user_id := c.Query("user_id")
	id, _ := strconv.Atoi(user_id)
	plan, err := db.FindPlanByUserID(id)
	if plan == nil {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
			"Plan":       nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "",
		"Plan":       plan,
	})
}
