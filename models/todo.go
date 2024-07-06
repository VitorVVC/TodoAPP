package models

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	InProgress  bool   `json:"in_progress"`
	Priority    bool   `json:"priority"`
}
