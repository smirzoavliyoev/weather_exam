package responser

import (
	"encoding/json"
	"io"
	"net/http"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(resp response, w http.ResponseWriter) error {
	r, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	_, err = io.WriteString(w, string(r))

	if err != nil {
		return err
	}

	return nil
}

var (
	InternalError = response{Code: 500, Message: "internal error"}
	BadRequest    = response{Code: 400, Message: "bad request"}
)
