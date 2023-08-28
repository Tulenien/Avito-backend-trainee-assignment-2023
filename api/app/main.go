package main

import (
    "log"
    "os"

    gin "github.com/gin-gonic/gin"
)

func someFunc(val int) int {
	return val + 2
}

func main() {
 router := gin.Default()

 router.GET("/ping", func(c *gin.Context) {
  c.JSON(200, gin.H{
   "message": "pong",
  })
 })

 router.Run(":80")
}