package repository

import (
	"context"
	"database/sql"

	"github.com/jevitapearl/TaskForge/internal/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func (pr *PostgresRepository) ExistsByTitle(ctx context.Context, title string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM tasks WHERE title=$1)`
	var exists bool

	pr.db.QueryRowContext(ctx, query, title).Scan(&exists)
	return exists
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (pr *PostgresRepository) GetAll(ctx context.Context) ([]models.Task, error) {

	query := `SELECT task_id, title, completed FROM tasks;`
	rows, err := pr.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	response := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
			return nil, err
		}
		response = append(response, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return response, nil
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

	rows, err := pr.db.ExecContext(ctx, query, new.Title, new.Completed, id)
	rowsAffected, _ := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (pr *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE task_id=$1`

	rows, err := pr.db.ExecContext(ctx, query, id)
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (pr *PostgresRepository) Close() error {
	return pr.db.Close()
}
