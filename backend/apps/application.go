package apps

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/shunyaYoshimra/checkCORS/routes"
	"github.com/shunyaYoshimra/checkCORS/database"
)

type Application struct{}

func (a Application) CreateApp(r *gin.Engine) {
	ConfigureAppDB()
	ConfigureAPIEndpoint(r)
	ConfigureView(r)
}

func (a Application) CreateProductionApp(r *gin.Engine) {
	ConfigureAPIEndpoint(r)
	ConfigureView(r)
}

func ConfigureAppDB() {
	database.AppConnaection()
	conn := database.GetDB()
	database.Migrate(conn)
}

func ConfigureProductionDB() {
	database.ProductionAppConnection()
	conn := database.GetDB()
	database.Migrate(conn)
}

func ConfigureAPIEndpoint(r *gin.Engine) {
	g := r.Group("/api/")
	routes.NewRouter(g)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "The page you are looking for dosen't exist",
		})
	})
}

func ConfigureView(r *gin.Engine) {
	r.Static("/app", "../dist")
}