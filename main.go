package main

import (
	"github.com/actionX/api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Connect()
	initRouter(r)
	r.Run()
}
