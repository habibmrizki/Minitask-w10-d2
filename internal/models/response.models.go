package models

type Response struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status,omitempty"`
	Name    string `json:"name,omitempty"`
}
