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

func (pr *PostgresRepository) GetAll(ctx context.Context, userID string) ([]models.Task, error) {

	query := `SELECT task_id, title, status FROM tasks WHERE user_id=$1;`
	rows, err := pr.db.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	response := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			return nil, err
		}
		response = append(response, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return response, nil
}

func (pr *PostgresRepository) GetByID(ctx context.Context, userID string, id string) (models.Task, error) {
	query := `SELECT task_id, title, completed FROM tasks WHERE task_id=$1 AND user_id=$2`

	var response models.Task
	if err := pr.db.QueryRowContext(ctx, query, id, userID).Scan(&response.ID, &response.Title, &response.Status); err != nil {
		return models.Task{}, err
	}
	return response, nil
}

func (pr *PostgresRepository) Create(ctx context.Context, userID string, task models.Task) error {
	query := `INSERT INTO tasks(title, status, user_id) VALUES($1, $2, $3)`

	if _, err := pr.db.ExecContext(ctx, query, task.Title, task.Status, userID); err != nil {
		return err
	}
	return nil
}

func (pr *PostgresRepository) Update(ctx context.Context, userID string, id string, new models.Task) error {
	query := `UPDATE tasks SET title=$1, status=$2 WHERE task_id=$3 AND user_id=$4`

	rows, err := pr.db.ExecContext(ctx, query, new.Title, new.Status, id, userID)
	rowsAffected, _ := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (pr *PostgresRepository) Delete(ctx context.Context, userID string, id string) error {
	query := `DELETE FROM tasks WHERE task_id=$1 AND user_id=$2`

	rows, err := pr.db.ExecContext(ctx, query, id, userID)
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (pr *PostgresRepository) Close() error {
	return pr.db.Close()
}
