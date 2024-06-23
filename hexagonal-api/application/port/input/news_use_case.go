package input

import (
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/application/domain"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/configuration/rest_err"
)

type NewsUseCase interface {
	GetNews(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
