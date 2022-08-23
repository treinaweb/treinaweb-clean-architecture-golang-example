package controllers

import (
	"github.com/cleysonph/clean-arch-go/infra/web"
)

type Controller interface {
	Execute(request web.HttpRequest) web.HttpResponse
}
