package models

type CreateTodoRequest struct {
	TodoName string `json:"todo_name" validate:"required"`
	IsCheck  bool   `json:"is_check"`
}

type ResponseReadTodo struct {
	ID       int64  `json:"id"`
	TodoName string `json:"todo_name"`
	IsCheck  bool   `json:"is_check"`
}

type UpdateTodoRequest struct {
	ID       int64  `json:"id" validate:"required"`
	TodoName string `json:"todo_name" validate:"required"`
	IsCheck  bool   `json:"is_check"`
}

type DeleteTodoRequest struct {
	ID int64 `json:"id" validate:"required"`
}

type ReadTodoRequest struct {
	IsCheck *bool `json:"is_check"`
}
