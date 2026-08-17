package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/jevitapearl/TaskForge/internal/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func New() *PostgresRepository {
	db, err := sql.Open("postgres", "host=localhost dbname=taskforge user=postgres password=password sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &PostgresRepository{db: db}
}

func (pr *PostgresRepository) GetAll(ctx context.Context) []models.Task {
	query := `SELECT task_id, title, completed FROM tasks;`
	rows, err := pr.db.QueryContext(ctx, query)
	if err != nil {
		return []models.Task{}
	}
	defer rows.Close()

	response := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
			return []models.Task{}
		}
		response = append(response, task)
	}

	if err := rows.Err(); err != nil {
		return []models.Task{}
	}
	return response
}

func (pr *PostgresRepository) GetByID(ctx context.Context, id string) (models.Task, error) {
	query := `SELECT task_id, title, completed FROM tasks WHERE task_id=$1`

	var response models.Task
	if err := pr.db.QueryRowContext(ctx, query, id).Scan(&response.ID, &response.Title, &response.Completed); err != nil {
		return models.Task{}, err
	}
	return response, nil
}

func (pr *PostgresRepository) Create(ctx context.Context, task models.Task) error {
	query := `INSERT INTO tasks(title, completed) VALUES($1, $2)`

	if _, err := pr.db.ExecContext(ctx, query, task.Title, task.Completed); err != nil {
		return err
	}
	return nil
}

func (pr *PostgresRepository) Update(ctx context.Context, id string, new models.Task) error {
	query := `UPDATE tasks SET title=$1, completed=$2 WHERE task_id=$3`

	if err := pr.db.QueryRowContext(ctx, query, new.Title, new.Completed, id); err != nil {
		return sql.ErrNoRows
	}
	return nil
}

func (pr *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE task_id=$1`

	if _, err := pr.db.ExecContext(ctx, query, id); err != nil {
		return sql.ErrNoRows
	}
	return nil
}

func (pr *PostgresRepository) Close() error {
	return pr.db.Close()
}
