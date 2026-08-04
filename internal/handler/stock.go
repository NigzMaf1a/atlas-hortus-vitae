package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/stock"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/queries"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func CreateStock(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		s, err := scripts.DecodeJSON[stock.Stock](r.Body)

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

		s, err = stock.CreateStock(
			ctx,
			db,
			queries.StockQueries.CreateStock,
			s,
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

		if err := scripts.EncodeJSON(w, s); err != nil {
			return
		}
	}
}

func ReadStock(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		stocks, err := stock.ReadStock(
			ctx,
			db,
			queries.StockQueries.ReadStock,
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

		if err := scripts.EncodeJSON(w, stocks); err != nil {
			return
		}
	}
}

func ReadOutletStock(db *sql.DB) http.HandlerFunc {
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

		stocks, err := stock.ReadOutletStock(
			ctx,
			db,
			queries.StockQueries.ReadOutletStock,
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

		if err := scripts.EncodeJSON(w, stocks); err != nil {
			return
		}
	}
}

func UpdateStockQty(db *sql.DB) http.HandlerFunc {
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

		qty, err := scripts.DecodeJSON[float64](
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

		err = stock.UpdateStockQty(
			ctx,
			db,
			queries.StockQueries.UpdateStockQty,
			qty,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Stock not found",
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

func UpdateStockPrice(db *sql.DB) http.HandlerFunc {
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

		cost, err := scripts.DecodeJSON[float64](
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

		err = stock.UpdateStockPrice(
			ctx,
			db,
			queries.StockQueries.UpdateStockPrice,
			cost,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Stock not found",
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