package repository

import (
	"database/sql"

	"github.com/Desgue/blog-tutorials/todo-api/domain"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	query := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRow(query, todo.Title, todo.Description, todo.Completed).Scan(&todo.Id, &todo.CreatedAt)
}

func (r *TodoRepository) GetByID(id int) (*domain.Todo, error) {
	todo := &domain.Todo{}
	query := `SELECT id, title, description, completed, created_at FROM todos WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) Update(todo *domain.Todo) error {
	query := `UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4`
	_, err := r.db.Exec(query, todo.Title, todo.Description, todo.Completed, todo.Id)
	return err
}

func (r *TodoRepository) Delete(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *TodoRepository) GetAll() ([]*domain.Todo, error) {
	query := `SELECT id, title, description, completed, created_at FROM todos`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*domain.Todo
	for rows.Next() {
		todo := &domain.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
