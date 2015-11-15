package models

type (
	Error struct {
		Error  string `json:"error"`
		Detail string `json:"detail"`
	}
)
