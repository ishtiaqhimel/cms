package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
)

var (
	ErrUnauthorized        = errors.New("role is unauthorized to perform this action")
	ErrNotFound            = errors.New("resource not found")
	ErrInvalidPage         = errors.New("invalid page request")
	ErrConflict            = errors.New("data conflict or already exist")
	ErrBadRequest          = errors.New("bad request, check param or body")
	ErrInternalServerError = errors.New("internal server error")
)

func getStatusCode(err error) int {
	switch err {
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInvalidPage:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrInternalServerError:
		return http.StatusInternalServerError
	default:
		wrapErr := &WrapErr{}
		if errors.As(err, wrapErr) {
			return wrapErr.StatusCode
		}
		return http.StatusInternalServerError
	}
}

// RespondError takes an `error` and a `customErr message` args
// to log the error to system and return to client
func RespondError(err error, customErr ...error) (int, Response) {
	var combinedErr error
	resp := Response{Success: false, Message: err.Error()}
	if len(customErr) > 0 {
		resp.Message = customErr[0].Error()
		combinedErr = fmt.Errorf("%s : %s", err, customErr[0])
	}
	statusCode := getStatusCode(err)
	if statusCode == http.StatusInternalServerError && config.Get().App.Disable500ErrMsgInResponse {
		resp.Message = ErrInternalServerError.Error()
	}
	logrus.Errorln(combinedErr)
	return statusCode, resp
}

type WrapErr struct {
	StatusCode int
	ErrCode    string
	Err        error
}

// implements error interface
func (e WrapErr) Error() string {
	return e.Err.Error()
}
