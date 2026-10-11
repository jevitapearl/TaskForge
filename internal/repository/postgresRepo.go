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

<<<<<<< HEAD
func (pr *PostgresRepository) ExistsByTitle(
	ctx context.Context,
	userID string,
	title string,
) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM tasks
			WHERE user_id = $1
			AND title = $2
		)
	`

	var exists bool

	err := pr.db.QueryRowContext(
		ctx,
		query,
		userID,
		title,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
=======
func (pr *PostgresRepository) ExistsByTitle(ctx context.Context, title string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM tasks WHERE title=$1)`
	var exists bool

	pr.db.QueryRowContext(ctx, query, title).Scan(&exists)
	return exists
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

<<<<<<< HEAD
func (pr *PostgresRepository) GetAll(ctx context.Context, userID string, filter models.TaskFilter) ([]models.Task, error) {

	offset := filter.Page * filter.Limit

	query := `
		SELECT task_id, title, status
		FROM tasks
		WHERE user_id = $1
		AND ($2::boolean IS NULL OR status = $2)
		ORDER BY task_id
		LIMIT $3 OFFSET $4
	`

	rows, err := pr.db.QueryContext(
		ctx,
		query,
		userID,
		filter.Completed,
		filter.Limit,
		offset,
	)
=======
func (pr *PostgresRepository) GetAll(ctx context.Context, userID string) ([]models.Task, error) {

	query := `SELECT task_id, title, status FROM tasks WHERE user_id=$1;`
	rows, err := pr.db.QueryContext(ctx, query, userID)
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c

	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	defer rows.Close()

	tasks := make([]models.Task, 0)
=======

	defer rows.Close()

	response := make([]models.Task, 0)
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c

	for rows.Next() {
		var task models.Task

<<<<<<< HEAD
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Status,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
=======
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			return nil, err
		}
		response = append(response, task)
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
<<<<<<< HEAD

	return tasks, nil
}

func (pr *PostgresRepository) GetByID(
	ctx context.Context,
	userID string,
	id string,
) (models.Task, error) {

	query := `
		SELECT task_id, title, status
		FROM tasks
		WHERE task_id = $1
		AND user_id = $2
	`

	var task models.Task

	err := pr.db.QueryRowContext(
		ctx,
		query,
		id,
		userID,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Status,
	)

	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (pr *PostgresRepository) Create(
	ctx context.Context,
	userID string,
	task models.Task,
) error {

	query := `
		INSERT INTO tasks(title, status, user_id)
		VALUES ($1, $2, $3)
	`

	_, err := pr.db.ExecContext(
		ctx,
		query,
		task.Title,
		task.Status,
		userID,
	)

	return err
}

func (pr *PostgresRepository) Update(ctx context.Context, userID string, id string, task models.Task) error {

	query := `UPDATE tasks SET title = $1, status = $2 WHERE task_id = $3 AND user_id = $4`

	result, err := pr.db.ExecContext(ctx, query, task.Title, task.Status, id, userID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (pr *PostgresRepository) Delete(
	ctx context.Context,
	userID string,
	id string,
) error {

	query := `
		DELETE FROM tasks
		WHERE task_id = $1 AND user_id = $2
	`

	result, err := pr.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
=======
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
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
