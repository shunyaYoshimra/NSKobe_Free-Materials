package controllers

import (
	"github.com/shunyaYoshimra/checkCORS/database"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/checkCORS/repositories"
	"github.com/shunyaYoshimra/checkCORS/middleware"
	// "strings"
)

var awsS3 *middleware.AwsS3
 

type MaterialController struct {
	Repository repositories.MaterialRepository
}

func NewMaterialController() MaterialController {
	return MaterialController{
		Repository: repositories.NewMaterialRepository(),
	}
}

func (mc *MaterialController) Index(c *gin.Context) {
	materials := mc.Repository.RetrieveMaterials()
	c.JSON(http.StatusOK, materials)
}

func (mc *MaterialController) Create(c *gin.Context) {
	var url string
	form, _ := c.MultipartForm()
	tags := c.Param("tags")
	time := c.PostForm("time")
	extension := c.PostForm("extension")
	fileHeaders := form.File["img"]
	if fileHeaders != nil {
		for _, fileHeader := range fileHeaders {
			awsS3 = middleware.NewAwsS3()
			file, err := fileHeader.Open()
			url, err = awsS3.Upload(file, time, extension)
			if err != nil {
					fmt.Print(err.Error())
					return
    	}
		}
	} else {
		return
	}
	material := database.Material{
		FileName: time + "." + extension, 
		Time: time,
		URL: url,
		Tags: tags,
	}
	if err := mc.Repository.Create(&material); err != nil {
		c.JSON(http.StatusBadRequest, "Something wrong happened")
	} else {
		c.Redirect(http.StatusFound, "/app/#/materials")
	}
}

func (mc *MaterialController) Delete(c *gin.Context) {
}

// http://localhost:3000/app/#/materials
// http://localhost:3000/app/#/materials