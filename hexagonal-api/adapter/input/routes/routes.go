package routes

import (
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/adapter/input/controller"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/adapter/output/news_http"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}
