package custom_errors

import (
	"auth-and-gateway-microservice/utils"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

type TranslatedError struct {
	Code int
	Err  error
}

func (t *TranslatedError) Error() string {
	return t.Err.Error()
}

func NewTranslatedError(code int, err error) *TranslatedError {
	return &TranslatedError{Code: code, Err: err}
}

func TranslateError(err error) *TranslatedError {
	var jwtErr utils.JwtError
	if errors.As(err, &jwtErr) {
		return NewTranslatedError(jwtErr.Code, jwtErr.Err)
	}
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		// postgres specific custom_errors
		switch err.(*pq.Error).Code {
		case "23505":
			return NewTranslatedError(http.StatusBadRequest, errors.New("key already exists"))
		case "22023":
			return NewTranslatedError(http.StatusBadRequest, errors.New("key already exists"))
		default:
			return NewTranslatedError(http.StatusInternalServerError, errors.New(strings.TrimPrefix(err.Error(), "pq: ")))
		}
	}
	return NewTranslatedError(0, fmt.Errorf("unknown error"))
}
