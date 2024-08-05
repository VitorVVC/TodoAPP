package models

import "github.com/google/uuid"

type Todo struct {
	BaseModel
	UUID        uuid.UUID `json:"uuid" query:"uuid"`
	Title       string    `json:"title" query:"title" validate:"required,max=30"`
	Description string    `json:"description" query:"description"`
	Done        bool      `json:"done" query:"done"`
	InProgress  bool      `json:"in_progress" query:"in_progress"`
	Priority    bool      `json:"priority" query:"priority"`
}
