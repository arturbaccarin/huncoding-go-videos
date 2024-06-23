package service

import (
	"fmt"

	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/application/domain"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/application/port/output"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/configuration/logger"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/configuration/rest_err"
)

type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNews(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("init getNews function in services, subject: %s, from: %s", newsDomain.Subject, newsDomain.From))

	newsDomainResponse, err := ns.newsPort.GetNewsPort(newsDomain)

	return newsDomainResponse, err
}
