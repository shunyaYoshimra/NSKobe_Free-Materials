package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("hello go")
	engine:= gin.Default()
    engine.GET("/api/message", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello Go",
        })
    })
    engine.Run(":3000")
}