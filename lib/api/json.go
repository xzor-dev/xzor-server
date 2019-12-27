package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse is the standard JSON response for API calls.
type JSONResponse struct {
	Status  Status      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

// JSONResponderFunc is a function used by JSONResponder to generate the data written to
// the http response writer.
type JSONResponderFunc func(w http.ResponseWriter, r *http.Request) *JSONResponse

// JSONResponder is a wrapper function used to send standard JSON responses.
func JSONResponder(f JSONResponderFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := f(w, r)
		code := http.StatusOK
		switch res.Status {
		case StatusError:
			code = http.StatusInternalServerError
		}
		b, err := json.Marshal(res)
		if err != nil {
			b = []byte(err.Error())
			code = http.StatusInternalServerError
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		_, err = w.Write(b)
		if err != nil {
			log.Fatalf("failed to write response: %v", err)
		}
	}
}

// NewJSONResponse creates a successfull JSONResponse with the supplied data.
func NewJSONResponse(data interface{}) *JSONResponse {
	return &JSONResponse{
		Status: StatusSuccess,
		Data:   data,
	}
}

// NewJSONErrorResponse creates a JSONResponse from an error.
func NewJSONErrorResponse(err error) *JSONResponse {
	return &JSONResponse{
		Status:  StatusError,
		Message: err.Error(),
	}
}

// NewJSONFailResponse creates a failure JSONResponse with the supplied message and data.
func NewJSONFailResponse(msg *string, data interface{}) *JSONResponse {
	res := &JSONResponse{
		Status: StatusFail,
		Data:   data,
	}
	if msg != nil {
		res.Message = *msg
	}
	return res
}
