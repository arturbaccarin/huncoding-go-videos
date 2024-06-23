package output

import (
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/application/domain"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/configuration/rest_err"
)

type NewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
