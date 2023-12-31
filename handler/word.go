package handler

import (
	"strconv"

	"github.com/actionX/api/db"
	"github.com/gin-gonic/gin"
)

func GetWordToday(c *gin.Context) {
	n_learn := c.Query("n_learn")
	learn_type := c.Query("type")
	dict_id := c.Query("dict_id")
	learn, _ := strconv.Atoi(n_learn)
	dict, _ := strconv.Atoi(dict_id)
	t, _ := strconv.Atoi(learn_type)
	var words []*db.Word
	var err error
	if t == 1 {
		words, err = db.FindWordByAlpha(dict, learn)

	} else if t == 2 {
		words, err = db.FindWordByAlphaDesc(dict, learn)
	} else if t == 3 {
		words, err = db.FindWordByRandom(dict, learn)
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
