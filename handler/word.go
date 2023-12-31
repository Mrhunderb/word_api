package handler

import (
	"strconv"

	"github.com/actionX/api/db"
	"github.com/gin-gonic/gin"
)

func GetWordToday(c *gin.Context) {
	var words []*db.Word
	var err error
	plan_id := c.Query("plan_id")
	plan, err := db.FindPlan(plan_id)
	if err != nil {
		c.JSON(200, gin.H{
			"Words": nil,
		})
		return
	}
	if plan.Mode == 1 {
		words, err = db.FindWordByAlpha(int(plan.DictID), int(plan.NLearn), int(plan.Progress))
	} else if plan.Mode == 2 {
		words, err = db.FindWordByAlphaDesc(int(plan.DictID), int(plan.NLearn), int(plan.Progress))
	} else if plan.Mode == 3 {
		words, err = db.FindWordByRandom(int(plan.DictID), int(plan.NLearn), int(plan.Progress))
	}
	if err != nil {
		c.JSON(200, gin.H{
			"Words": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"Words": words,
	})
}

func GetAllWord(c *gin.Context) {
	var words []*db.Word
	var err error
	dict_id := c.Query("dict_id")
	offset := c.Query("offset")
	n_offset, _ := strconv.Atoi(offset)
	words, err = db.FindAllWord(dict_id, n_offset)
	if err != nil {
		c.JSON(200, gin.H{
			"Words": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"Words": words,
	})
}
