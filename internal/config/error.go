package config

import (
	"errors"
	"net/http"

	"github.com/bencoronard/demo-go-common-libs/dto"
	xhttp "github.com/bencoronard/demo-go-common-libs/http"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
)

type appErrorHandler struct{}

func NewAppErrorHandler() xhttp.AppErrorHandler {
	return &appErrorHandler{}
}

func (a *appErrorHandler) Handle(err error, pd dto.ProblemDetail) (dto.ProblemDetail, bool) {
	switch {
	case errors.Is(err, resource.ErrResourceNotFound):
		return pd.
			WithStatus(http.StatusNotFound).
			WithDetail(err.Error()), true
	case errors.Is(err, resource.ErrOptimisticLockFailure):
		return pd.
			WithStatus(http.StatusConflict).
			WithDetail(err.Error()), true
	default:
		return pd, false
	}
}
