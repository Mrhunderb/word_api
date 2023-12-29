package main

import (
	"github.com/actionX/com.word.api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Connect()
	r.Run()
}
