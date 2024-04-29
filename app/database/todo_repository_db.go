package database

import (
	"context"

	"todo/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, createTodoRequest *models.CreateTodoRequest) error
	ReadTodo(ctx context.Context) (*[]models.ResponseReadTodo, error)
}

type TodoRepositoryDB struct {
	pool *pgxpool.Pool
}

func NewTodoRepositoryDB(pool *pgxpool.Pool) TodoRepository {
	return &TodoRepositoryDB{
		pool: pool,
	}
}

func (r *TodoRepositoryDB) CreateTodo(ctx context.Context, createTodoRequest *models.CreateTodoRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit(ctx)
		default:
			_ = tx.Rollback(ctx)
		}
	}()

	stmt := `INSERT INTO todo_list (todo_name,is_check)
        VALUES(@todo_name, @is_check);`
	args := pgx.NamedArgs{
		"todo_name": createTodoRequest.TodoName,
		"is_check":  createTodoRequest.IsCheck,
	}

	_, err = tx.Exec(ctx, stmt, args)
	if err != nil {
		return err
	}

	return err
}

func (r *TodoRepositoryDB) ReadTodo(ctx context.Context) (*[]models.ResponseReadTodo, error) {
	query := "SELECT tl.id, tl.todo_name, tl.is_check FROM todo_list tl"

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responseReadTodoList []models.ResponseReadTodo
	for rows.Next() {
		var responseReadTodo models.ResponseReadTodo
		err := rows.Scan(
			&responseReadTodo.ID,
			&responseReadTodo.TodoName,
			&responseReadTodo.IsCheck,
		)
		if err != nil {
			return nil, err
		}
		responseReadTodoList = append(responseReadTodoList, responseReadTodo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(responseReadTodoList) == 0 {
		return &[]models.ResponseReadTodo{}, nil
	}

	return &responseReadTodoList, nil
}