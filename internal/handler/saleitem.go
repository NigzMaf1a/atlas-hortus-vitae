package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/saleitem"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/queries"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func CreateSaleItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		item, err := scripts.DecodeJSON[saleitem.SaleItem](r.Body)

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

		item, err = saleitem.CreateSaleItem(
			ctx,
			db,
			queries.SaleItemQueries.CreateSaleItem,
			item,
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

		if err := scripts.EncodeJSON(w, item); err != nil {
			return
		}
	}
}

func ReadSaleItems(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		items, err := saleitem.ReadSaleItems(
			ctx,
			db,
			queries.SaleItemQueries.ReadSaleItems,
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

		if err := scripts.EncodeJSON(w, items); err != nil {
			return
		}
	}
}

func ReadSaleItemsBySale(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		saleID, err := scripts.ConvertToInteger(
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

		items, err := saleitem.ReadSaleItemsBySale(
			ctx,
			db,
			queries.SaleItemQueries.ReadSaleItemsBySale,
			saleID,
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

		if err := scripts.EncodeJSON(w, items); err != nil {
			return
		}
	}
}

func ReadSaleItemsByProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productID, err := scripts.ConvertToInteger(
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

		items, err := saleitem.ReadSaleItemsByProduct(
			ctx,
			db,
			queries.SaleItemQueries.ReadSaleItemsByProduct,
			productID,
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

		if err := scripts.EncodeJSON(w, items); err != nil {
			return
		}
	}
}
