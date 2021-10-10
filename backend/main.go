package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/checkCORS/apps"
)

func main() {
	r:= gin.Default()
	app := new(apps.Application)
	app.CreateApp(r)
	r.Run(":3000")
}

