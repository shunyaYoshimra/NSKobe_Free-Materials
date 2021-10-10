package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/checkCORS/controllers"
)

func NewRouter(g *gin.RouterGroup) {
	sampleController := controllers.NewSampleController()
	materialController := controllers.NewMaterialController()
	{
		g.GET("/sample", sampleController.SayHello)
		g.GET("/materials", materialController.Index)
		g.POST("/materials/:tags", materialController.Create)
		g.DELETE("/materials", materialController.Delete)
	}
}