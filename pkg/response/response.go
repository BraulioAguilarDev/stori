package response

import (
	"encoding/json"
	"log"
)

type Response struct {
	Success bool   `json:"success"`
	Errors  string `json:"errors,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success(data any) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func Failure(err string) Response {
	return Response{
		Success: false,
		Errors:  err,
	}
}

func FailureMappingErrors(errs map[string]string) Response {
	var msj string
	if len(errs) > 0 {
		b, e := json.Marshal(errs)
		if e != nil {
			log.Printf("Marshal error: %s\n", e.Error())
		}

		msj = string(b)
	}

	return Response{
		Success: false,
		Errors:  string(msj),
	}
}
