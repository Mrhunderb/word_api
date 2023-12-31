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

	apiRouter.GET("/dict/list/", handler.DictList)
	apiRouter.GET("/dict/", handler.GetDict)

	apiRouter.GET("/plan/", handler.GetPlan)

	apiRouter.GET("/word/today/", handler.GetWordToday)
	apiRouter.GET("/word/all/", handler.GetAllWord)
}
