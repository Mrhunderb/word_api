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
	plan_id_int, _ := strconv.Atoi(plan_id)
	plan, err := db.FindPlan(plan_id)
	if err != nil {
		c.JSON(200, gin.H{
			"Words": nil,
		})
		return
	}
	n, _ := db.GetTodyLearn(plan_id_int)
	if plan.Mode == 1 {
		words, err = db.FindWordByAlpha(int(plan.DictID), int(plan.NLearn), int(plan.Progress)-n)
	} else if plan.Mode == 2 {
		words, err = db.FindWordByAlphaDesc(int(plan.DictID), int(plan.NLearn), int(plan.Progress)-n)
	} else if plan.Mode == 3 {
		words, err = db.FindWordByRandom(int(plan.DictID), int(plan.NLearn), int(plan.Progress)-n)
	}
	if err != nil {
		c.JSON(200, gin.H{
			"Words":  nil,
			"Review": nil,
		})
		return
	}
	review, err := db.GetReviwWord(plan_id_int, int(plan.NReview))
	if err != nil {
		c.JSON(200, gin.H{
			"Words":  words,
			"Review": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"Words":  words,
		"Review": review,
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

func GetCollectWord(c *gin.Context) {
	var words []*db.Word
	var err error
	user_id := c.Query("user_id")
	words, err = db.FindUserCollection(user_id)
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

func AddCollectWord(c *gin.Context) {
	var err error
	user := c.Query("user_id")
	word := c.Query("word_id")
	user_id, _ := strconv.Atoi(user)
	word_id, _ := strconv.Atoi(word)
	err = db.AddUserCollection(user_id, word_id)
	if err != nil {
		c.JSON(200, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "添加成功",
	})
}

func DeletCollect(c *gin.Context) {
	var err error
	user := c.Query("user_id")
	word := c.Query("word_id")
	user_id, _ := strconv.Atoi(user)
	word_id, _ := strconv.Atoi(word)
	err = db.DeletCollection(user_id, word_id)
	if err != nil {
		c.JSON(200, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "删除成功",
	})
}

func AddHistory(c *gin.Context) {
	var err error
	plan_id := c.Query("plan_id")
	word_id := c.Query("word_id")
	is_know := c.Query("is_know")
	plan_id_int, _ := strconv.Atoi(plan_id)
	word_id_int, _ := strconv.Atoi(word_id)
	is_know_int, _ := strconv.Atoi(is_know)
	err = db.AddUserHistory(plan_id_int, word_id_int, is_know_int)
	if err != nil {
		c.JSON(200, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "添加成功",
	})
}

func GetTodyLearn(c *gin.Context) {
	var err error
	plan_id := c.Query("plan_id")
	plan_id_int, _ := strconv.Atoi(plan_id)
	n_learn, _ := db.GetTodyLearn(plan_id_int)
	n_review, err := db.GetTodyReview(plan_id_int)
	if err != nil {
		c.JSON(200, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "",
		"n_learn":    n_learn,
		"n_review":   n_review,
	})
}
