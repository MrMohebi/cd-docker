package router

import (
	"cd-docker/controllers"
	"github.com/gin-gonic/gin"
)

func Routs(r *gin.Engine) {
	AssetsRoute(r)
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", controllers.Index())
	r.GET("/docs", controllers.Docs())

	r.POST("/updateServiceLatest", controllers.UpdateServiceLatest())
	r.POST("/updateServiceWithImage", controllers.UpdateServiceWithImage())
}
