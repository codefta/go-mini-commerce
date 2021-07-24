package utils

import (
	"encoding/json"
	"net/http"
)

type APIResp struct {
	StatusCode	int			`json:"-"`
	Success		bool		`json:"success"`
	Data		interface{}	`json:"data,omitempty"`
	Err			string		`json:"err,omitempty"`
	Message		string		`json:"message,omitempty"`
}

func NewSuccessResp(statusCode int, data interface{}) APIResp {
	return APIResp{
		StatusCode: statusCode,
		Success: true,
		Data: data,
	}
}

func NewErrorResp(statusCode int, errCode string, message string) APIResp {
	return APIResp{
		StatusCode: statusCode,
		Err: errCode, 
		Message: message,
	}
}

func WriteResp(w http.ResponseWriter, resp APIResp) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}