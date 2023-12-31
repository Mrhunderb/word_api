package handler

import (
	"github.com/actionX/api/db"
	"github.com/gin-gonic/gin"
)

func DictList(c *gin.Context) {
	dict_list := db.GetDictList()
	c.JSON(200, gin.H{
		"Dicts": dict_list,
	})
}
