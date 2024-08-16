package repository

import (
	"database/sql"

	"github.com/Desgue/blog-tutorials/todo-api/domain"
	_ "github.com/lib/pq"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo domain.Todo) (*domain.Todo, error) {
	var createdTodo domain.Todo
	query := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING *`
	err := r.db.QueryRow(
		query,
		todo.Title,
		todo.Description,
		todo.Completed).
		Scan(
			&createdTodo.Id,
			&createdTodo.Title,
			&createdTodo.Description,
			&createdTodo.Completed,
			&createdTodo.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &createdTodo, nil
}

func (r *TodoRepository) GetByID(id int) (*domain.Todo, error) {
	todo := &domain.Todo{}
	query := `SELECT id, title, description, completed  FROM todos WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) Update(todo domain.Todo) (*domain.Todo, error) {
	var updatedTodo domain.Todo
	query := `UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4 RETURNING *`
	err := r.db.QueryRow(
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		todo.Id).
		Scan(
			&updatedTodo.Id,
			&updatedTodo.Title,
			&updatedTodo.Description,
			&updatedTodo.Completed,
			&updatedTodo.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &updatedTodo, nil
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
