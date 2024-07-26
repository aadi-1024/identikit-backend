package models

type JsonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
