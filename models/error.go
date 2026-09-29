package models

import (
	"encoding/json"
	"net/http"
)

type ErrorDTO struct {
	Message error
}

func (e *ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")

	if err != nil {
		panic(err)
	}

	return string(b)
}

func DoError(e error, w http.ResponseWriter, status ...int) {
	httpStatus := http.StatusBadRequest
	if len(status) > 0 {
		httpStatus = status[0]
	}

	errDTO := ErrorDTO{Message: e}
	http.Error(w, errDTO.ToString(), httpStatus)
}
