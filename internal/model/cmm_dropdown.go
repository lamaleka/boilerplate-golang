package model

type Dropdown[T any] struct {
	Key   *string `json:"key,omitempty"`
	Value T       `json:"value"`
	Label string  `json:"label"`
}
