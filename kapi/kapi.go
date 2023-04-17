package kapi

type Response[T any] struct {
	Error  []string `json:"error"`
	Result T        `json:"result,omitempty"`
}
