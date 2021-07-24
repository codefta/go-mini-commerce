package utils

import (
	"log"
	"net/http"
)

func NewBadRequestErrorResp(message string) APIResp {
	return NewErrorResp(http.StatusBadRequest, "ERR_BAD_REQUEST", message)
}

func NewInternalServerErrorResp(err error) APIResp {
	log.Println(err)
	return NewErrorResp(http.StatusInternalServerError, "ERR_INTERNAL_SERVER", "")
}

func NewNotFoundErrorResp(message string) APIResp {
	return NewErrorResp(http.StatusNotFound, "ERR_NOT_FOUND", message)
}