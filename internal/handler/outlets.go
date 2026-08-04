package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/outlets"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/queries"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func CreateOutlet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		outlet, err := scripts.DecodeJSON[outlets.Outlet](r.Body)

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

		outlet, err = outlets.CreateOutlet(
			ctx,
			db,
			queries.OutletQueries.CreateOutlet,
			outlet,
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

		if err := scripts.EncodeJSON(w, outlet); err != nil {
			return
		}
	}
}

func ReadOutlets(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		outlets, err := outlets.ReadOutlets(
			ctx,
			db,
			queries.OutletQueries.ReadOutlets,
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

		if err := scripts.EncodeJSON(w, outlets); err != nil {
			return
		}
	}
}

func ReadOutlet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := scripts.ConvertToInteger(
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

		outlet, err := outlets.ReadOutlet(
			ctx,
			db,
			queries.OutletQueries.ReadOutlet,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Outlet not found",
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

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		if err := scripts.EncodeJSON(w, outlet); err != nil {
			return
		}
	}
}

func UpdateNetworth(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		id, err := scripts.ConvertToInteger(
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

		networth, err := scripts.DecodeJSON[float64](
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

		err = outlets.UpdateNetworth(
			ctx,
			db,
			queries.OutletQueries.UpdateNetworth,
			networth,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Outlet not found",
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

func UpdateOutletStatus(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		id, err := scripts.ConvertToInteger(
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

		status, err := scripts.DecodeJSON[string](
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

		err = outlets.UpdateOutletStatus(
			ctx,
			db,
			queries.OutletQueries.UpdateOutletStatus,
			status,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Outlet not found",
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
