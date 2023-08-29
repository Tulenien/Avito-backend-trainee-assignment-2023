package main

import (
    gin "github.com/gin-gonic/gin"
)

func someFunc(val int) int {
	return val + 2
}

func main() int {
 router := gin.Default()

 router.GET("/ping", func(c *gin.Context) {
  c.JSON(200, gin.H{
   "message": "pong",
  })
 })
 
 err = router.Run(":80")
 if (err != nil) {
    return 0;
 }
 return 1;
}