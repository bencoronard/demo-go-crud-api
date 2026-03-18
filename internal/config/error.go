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

func (a *appErrorHandler) Handle(err error, pd *dto.ProblemDetail) error {
	switch {
	case errors.Is(err, resource.ErrResourceNotFound):
		pd.Status = http.StatusNotFound
		pd.Detail = err.Error()
		return nil
	case errors.Is(err, resource.ErrOptimisticLockFailure):
		pd.Status = http.StatusConflict
		pd.Detail = err.Error()
		return nil
	}
	return err
}
