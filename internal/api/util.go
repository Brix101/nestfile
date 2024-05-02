package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Brix101/nestfile/internal/domain"
	"github.com/Brix101/nestfile/internal/middlewares"
	"github.com/Brix101/nestfile/internal/util"
	"github.com/go-playground/validator"
	"github.com/mattn/go-sqlite3"
)

func (a api) responseJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	marsh, err := json.Marshal(data)
	if err != nil {
		a.responseError(w, r, err, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(marsh)
}

func (a api) responseError(w http.ResponseWriter, _ *http.Request, err error, status int) {
	var errData domain.ErrResponse

	switch typedErr := err.(type) {
	case sqlite3.Error:
		fieldName := "field" // Default field name if constraint name is not in the expected format
		parts := strings.Split(typedErr.Error(), ":")
		if len(parts) > 0 {
			// Extract the field name from the last part after the last dot
			lastPart := strings.TrimSpace(parts[len(parts)-1])
			// Find the last dot index
			dotIndex := strings.LastIndex(lastPart, ".")
			if dotIndex != -1 && dotIndex < len(lastPart)-1 {
				// Extract the field name after the last dot
				fieldName = strings.TrimSpace(lastPart[dotIndex+1:])
			}
		}

		switch typedErr.Code {
		case sqlite3.ErrConstraint:
			status = http.StatusBadRequest
			message := fmt.Sprintf("The %s you entered is already taken.", fieldName)
			errorFields := []domain.ErrField{{
				Field:   fieldName,
				Message: message,
			}}

			errData = domain.ErrResponse{
				Message: "Validation Error",
				Errors:  errorFields,
			}

		default:
			status = http.StatusInternalServerError
			errData = domain.ErrResponse{
				Message: "Something went wrong!",
				Errors:  []domain.ErrField{},
			}
		}

	case validator.ValidationErrors:
		errorFields := []domain.ErrField{}
		for _, e := range typedErr {
			errorFields = append(errorFields, domain.ErrField{
				Field:   strings.ToLower(e.Field()),
				Message: util.GetValidationErrorMessage(e),
			})
		}

		errData = domain.ErrResponse{
			Message: "Validation Error",
			Errors:  errorFields,
		}

	case error:
		var message string
		if status == 500 {
			message = "Something went wrong!"
		} else {
			message = typedErr.Error()
		}
		errData = domain.ErrResponse{
			Message: message,
			Errors:  []domain.ErrField{},
		}

	default:
		status = 500
		errData = domain.ErrResponse{
			Message: "Something went wrong!",
			Errors:  []domain.ErrField{},
		}

	}

	marsh, _ := json.Marshal(errData)

	w.Header().Set("X-Nestfile-Error", err.Error())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(marsh)
}

func (a api) getAuthUser(_ http.ResponseWriter, r *http.Request) *domain.User {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var user *domain.User
	usrCtx := ctx.Value(middlewares.UserCtxKey{})
	if usrCtx != nil {
		item, ok := usrCtx.(*domain.AuthToken)
		if ok {
			usr, err := a.userRepo.GetByID(ctx, int64(item.Sub))
			if err == nil {
				user = &usr
			}
		}
	}

	return user
}
