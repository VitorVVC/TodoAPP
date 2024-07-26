package models

type Todo struct {
	ID          int64  `json:"id" query:"id"`
	Title       string `json:"title" query:"title"`
	Description string `json:"description" query:"description"`
	Done        bool   `json:"done" query:"done"`
	InProgress  bool   `json:"inProgress" query:"inProgress"`
	Priority    bool   `json:"priority" query:"priority"`
}
