package handler

import (
	"net/http"
	"strconv"

	"github.com/actionX/api/db"
	"github.com/gin-gonic/gin"
)

func DictList(c *gin.Context) {
	dict_list := db.GetDictList()
	c.JSON(http.StatusOK, gin.H{
		"Dicts": dict_list,
	})
}

func GetDict(c *gin.Context) {
	dict_id := c.Query("dict_id")
	id, _ := strconv.Atoi(dict_id)
	dict, err := db.FindDict(id)
	if dict == nil {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
			"Dict":       nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "",
		"Dict":       dict,
	})
}
