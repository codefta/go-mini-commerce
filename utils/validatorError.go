package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)


type ValErrResp struct {
	Errors []string `json:"errors"`
}

func ValidatorError(err error) *ValErrResp {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ValErrResp{
			Errors: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("`%s` cannot be empty", err.Field())
			case "numeric":
				resp.Errors[i] = fmt.Sprintf("`%s` must numeric", err.Field())
			default:
				resp.Errors[i] = fmt.Sprintf("something wrong on %s %s", err.Field(), err.Tag())
			}
		}

		return &resp
	}
	return nil
}
