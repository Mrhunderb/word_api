package main

import (
	"github.com/actionX/api/handler"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("")

	// basic apis
	apiRouter.POST("/user/login/", handler.Login)
	apiRouter.POST("/user/register/", handler.Register)
	apiRouter.GET("/user/info/", handler.GetInfo)
	apiRouter.GET("/user/collect/", handler.GetCollectWord)
	apiRouter.GET("/user/plan/", handler.GetUserPlan)

	apiRouter.POST("/collect/add/", handler.AddCollectWord)
	apiRouter.POST("/collect/delete/", handler.DeletCollect)

	apiRouter.GET("/dict/list/", handler.DictList)
	apiRouter.GET("/dict/", handler.GetDict)

	apiRouter.GET("/plan/", handler.GetPlan)
	apiRouter.POST("/plan/change/", handler.ChangePlan)

	apiRouter.GET("/word/today/", handler.GetWordToday)
	apiRouter.GET("/word/today/learn", handler.GetTodyLearn)
	apiRouter.GET("/word/all/", handler.GetAllWord)
	apiRouter.POST("/word/history/", handler.AddHistory)
}
