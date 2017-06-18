package db

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
	UpdatedAt   string `json:"updatedAt"`
	CreatedAt   string `json:"createdAt"`
}
