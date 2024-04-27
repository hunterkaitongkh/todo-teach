package models

type CreateTodoRequestDataBase struct {
	TodoName string `db:"todo_name" json:"todo_name" `
	IsCheck  bool   `db:"is_check" json:"is_check" `
}

type CreateTodoRequest struct {
	TodoName string `json:"todo_name" validate:"required"`
	IsCheck  bool   `json:"is_check"`
}
