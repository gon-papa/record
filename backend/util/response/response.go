package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type errorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message,omitempty"`
	Trace      string `json:"trace,omitempty"`
}

var (
	ErrUnknown             = NewErrorResponse(400, "unknown", "")
	ErrUnAuthorized        = NewErrorResponse(401, "unauthorized", "")
	ErrRecordNotFound      = NewErrorResponse(404, "not found", "")
	ErrInternalServerError = NewErrorResponse(500, "internal server error", "")
)

type Responser interface {
	CreateResponse(w http.ResponseWriter, status int)
}

func (r *response) CreateResponse(w http.ResponseWriter, status int) {
	res, err := json.Marshal(&response{Message: r.Message, Data: r.Data})
	if err != nil {
		handleError(w, http.StatusInternalServerError, "internal server error", fmt.Sprintf("error marshaling response: %v", err))
		return
	}

	createResponse(w, status, res)
}

func (e *errorResponse) CreateResponse(w http.ResponseWriter, status int) {
	res, err := json.Marshal(&errorResponse{StatusCode: status, Message: e.Message, Trace: e.Trace})
	if err != nil {
		handleError(w, http.StatusInternalServerError, "internal server error", fmt.Sprintf("error marshaling error response: %v", err))
		return
	}

	createResponse(w, status, res)
}

func createResponse(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(body)
}

func handleError(w http.ResponseWriter, status int, message string, trace string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	err := &errorResponse{StatusCode: status, Message: message, Trace: trace}
	res, _ := json.Marshal(err)
	w.Write(res)

	// ログ出力
}

func NewResponse(msg string, data interface{}) *response {
	return &response{Message: msg, Data: data}
}

func NewErrorResponse(statusCode int, msg string, trace string) *errorResponse {
	return &errorResponse{StatusCode: statusCode, Message: msg, Trace: trace}
}

func (e *errorResponse) Error() string {
	return fmt.Sprintf("code: %d, message: %s, trace: %s", e.StatusCode, e.Message, e.Trace)
}
