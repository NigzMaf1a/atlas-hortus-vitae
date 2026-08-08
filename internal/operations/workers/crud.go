package workers

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func CreateWorker(
	ctx context.Context,
	db *sql.DB,
	query string,
	w Worker,
) (Worker, error) {

	err := db.QueryRowContext(
		ctx,
		query,
		w.UserId,
		w.OutletID,
		w.WorkerPresent,
		w.ShiftTime,
		w.SignInLocation,
	).Scan(
		&w.UserId,
		&w.OutletID,
		&w.WorkerPresent,
		&w.SignInTime,
		&w.ShiftTime,
		&w.SignInLocation,
	)

	if err != nil {
		return w, fmt.Errorf("create worker: %w", err)
	}

	fmt.Println("Worker created successfully")

	return w, nil
}

func GetWorkers(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Worker, error) {

	workers := []Worker{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("read workers: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var w Worker

		err := rows.Scan(
			&w.UserId,
			&w.OutletID,
			&w.WorkerPresent,
			&w.SignInTime,
			&w.ShiftTime,
			&w.SignInLocation,
		)

		if err != nil {
			return nil, fmt.Errorf("scan worker: %w", err)
		}

		workers = append(workers, w)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate workers: %w", err)
	}

	fmt.Println("Workers fetched successfully")

	return workers, nil
}

func GetWorkersByOutlet(
	ctx context.Context,
	db *sql.DB,
	query string,
	outletID int64,
) ([]Worker, error) {

	workers := []Worker{}

	rows, err := db.QueryContext(
		ctx,
		query,
		outletID,
	)

	if err != nil {
		return nil, fmt.Errorf("read workers by outlet: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var w Worker

		err := rows.Scan(
			&w.UserId,
			&w.OutletID,
			&w.WorkerPresent,
			&w.SignInTime,
			&w.ShiftTime,
			&w.SignInLocation,
		)

		if err != nil {
			return nil, fmt.Errorf("scan worker: %w", err)
		}

		workers = append(workers, w)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate outlet workers: %w", err)
	}

	fmt.Println("Outlet workers fetched successfully")

	return workers, nil
}

func SignInWorker(
	ctx context.Context,
	db *sql.DB,
	query string,
	userID int64,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return fmt.Errorf("sign in worker: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Worker signed in successfully")

	return nil
}

func UpdateShiftTime(
	ctx context.Context,
	db *sql.DB,
	query string,
	shiftTime time.Time,
	userID int64,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		shiftTime,
		userID,
	)

	if err != nil {
		return fmt.Errorf("update worker shift time: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Worker shift time updated successfully")

	return nil
}

func UpdateSignInLocation(
	ctx context.Context,
	db *sql.DB,
	query string,
	location string,
	userID int64,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		location,
		userID,
	)

	if err != nil {
		return fmt.Errorf("update worker sign-in location: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Worker sign-in location updated successfully")

	return nil
}

func UpdateOutletID(
	ctx context.Context,
	db *sql.DB,
	query string,
	outletID int64,
	userID int64,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		outletID,
		userID,
	)

	if err != nil {
		return fmt.Errorf("update worker outlet: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Worker outlet updated successfully")

	return nil
}
