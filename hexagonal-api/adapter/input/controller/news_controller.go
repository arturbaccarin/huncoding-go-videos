package controller

import (
	"net/http"

	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/adapter/input/model/request"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/application/domain"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/application/port/input"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/configuration/logger"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(newsUseCase input.NewsUseCase) *newsController {
	return &newsController{newsUseCase}
}

func (nc *newsController) GetNews(c *gin.Context) {
	logger.Info("Init GetNews controller api")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From:    newsRequest.From.Format("2006-01-02"),
	}

	newsResponseDomain, err := nc.newsUseCase.GetNews(newsDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsResponseDomain)
}
