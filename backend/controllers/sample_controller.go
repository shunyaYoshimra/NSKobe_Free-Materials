package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SampleController struct {}

func NewSampleController() SampleController {
	return SampleController{}
}

func (sc *SampleController) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello from sampleController")
}