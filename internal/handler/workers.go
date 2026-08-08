package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/workers"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/queries"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func CreateWorker(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		worker, err := scripts.DecodeJSON[workers.Worker](r.Body)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		worker, err = workers.CreateWorker(
			ctx,
			db,
			queries.WorkerQueries.CreateWorker,
			worker,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		w.WriteHeader(http.StatusCreated)

		if err := scripts.EncodeJSON(w, worker); err != nil {
			return
		}
	}
}

func GetWorkers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		workers, err := workers.GetWorkers(
			ctx,
			db,
			queries.WorkerQueries.GetWorkers,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		if err := scripts.EncodeJSON(w, workers); err != nil {
			return
		}
	}
}

func GetWorkersByOutlet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		outletID, err := scripts.ConvertToInteger(
			r.PathValue("id"),
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		workers, err := workers.GetWorkersByOutlet(
			ctx,
			db,
			queries.WorkerQueries.GetWorkersByOutlet,
			outletID,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		if err := scripts.EncodeJSON(w, workers); err != nil {
			return
		}
	}
}

func SignInWorker(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, err := scripts.ConvertToInteger(
			r.PathValue("id"),
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		err = workers.SignInWorker(
			ctx,
			db,
			queries.WorkerQueries.SignInWorker,
			userID,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Worker not found",
					http.StatusNotFound,
				)
				return
			}

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateShiftTime(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		userID, err := scripts.ConvertToInteger(
			r.PathValue("id"),
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		shiftTime, err := scripts.DecodeJSON[time.Time](
			r.Body,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		err = workers.UpdateShiftTime(
			ctx,
			db,
			queries.WorkerQueries.UpdateShiftTime,
			shiftTime,
			userID,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Worker not found",
					http.StatusNotFound,
				)
				return
			}

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateSignInLocation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		userID, err := scripts.ConvertToInteger(
			r.PathValue("id"),
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		location, err := scripts.DecodeJSON[string](
			r.Body,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		err = workers.UpdateSignInLocation(
			ctx,
			db,
			queries.WorkerQueries.UpdateSignInLocation,
			location,
			userID,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Worker not found",
					http.StatusNotFound,
				)
				return
			}

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateOutletID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		userID, err := scripts.ConvertToInteger(
			r.PathValue("id"),
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		outletID, err := scripts.DecodeJSON[int64](
			r.Body,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		err = workers.UpdateOutletID(
			ctx,
			db,
			queries.WorkerQueries.UpdateOutletID,
			outletID,
			userID,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Worker not found",
					http.StatusNotFound,
				)
				return
			}

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
