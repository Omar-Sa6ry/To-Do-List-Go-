package models

import (
	"time"
	"todolist/db"
)

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	UserID      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (t *Todo) Save() error {
	query := `
		INSERT INTO todos(title, description, status, user_id) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	
	if t.Status == "" {
		t.Status = "pending"
	}

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(t.Title, t.Description, t.Status, t.UserID).Scan(&t.ID, &t.CreatedAt)
	return err
}

func GetAllTodos(userId int64) ([]Todo, error) {
	query := "SELECT id, title, description, status, user_id, created_at FROM todos WHERE user_id = $1"
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	if todos == nil {
		todos = []Todo{}
	}
	return todos, nil
}

func GetTodoByID(id int64, userId int64) (*Todo, error) {
	query := "SELECT id, title, description, status, user_id, created_at FROM todos WHERE id = $1 AND user_id = $2"
	row := db.DB.QueryRow(query, id, userId)

	var t Todo
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Todo) Update() error {
	query := `
		UPDATE todos 
		SET title = $1, description = $2, status = $3 
		WHERE id = $4 AND user_id = $5`
	
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Title, t.Description, t.Status, t.ID, t.UserID)
	return err
}

func DeleteTodo(id int64, userId int64) error {
	query := "DELETE FROM todos WHERE id = $1 AND user_id = $2"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, userId)
	return err
}
