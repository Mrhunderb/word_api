package handler

import (
	"net/http"
	"strconv"

	"github.com/actionX/api/db"
	"github.com/gin-gonic/gin"
)

func GetUserPlan(c *gin.Context) {
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

func GetPlan(c *gin.Context) {
	user_id := c.Query("user_id")
	dict_id := c.Query("dict_id")

	plan, err := db.FindPlanByID(user_id, dict_id)
	if plan == nil {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
			"Plan": db.Plan{
				PlanID:   0,
				Mode:     0,
				NLearn:   30,
				NReview:  30,
				Progress: 0,
				DictID:   0,
				UserID:   0,
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "",
		"Plan":       plan,
	})
}

func ChangePlan(c *gin.Context) {
	user := c.Query("user_id")
	dict := c.Query("dict_id")
	mode := c.Query("mode")
	learn := c.Query("n_learn")
	review := c.Query("n_review")

	user_id, _ := strconv.Atoi(user)
	dict_id, _ := strconv.Atoi(dict)
	mode_id, _ := strconv.Atoi(mode)
	n_learn, _ := strconv.Atoi(learn)
	n_review, _ := strconv.Atoi(review)

	var plan *db.Plan
	plan, err := db.FindPlanByID(user, dict)

	if err != nil {
		plan, _ = db.AddPlan(user_id, dict_id, mode_id, n_learn, n_review)
		db.UpdateUserPlan(user_id, int(plan.PlanID))
		c.JSON(200, gin.H{
			"StatusCode": 0,
			"StatusMsg":  "",
			"PlanId":     0,
		})
		return
	}
	db.UpdatePlan(*plan, mode_id, n_learn, n_review)
	db.UpdateUserPlan(user_id, int(plan.PlanID))
	c.JSON(200, gin.H{
		"StatusCode": 0,
		"StatusMsg":  "",
		"PlanId":     plan.PlanID,
	})
}
