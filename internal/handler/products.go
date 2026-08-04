package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/products"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/queries"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		product, err := scripts.DecodeJSON[products.Product](r.Body)

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

		product, err = products.CreateProduct(
			ctx,
			db,
			queries.ProductQueries.CreateProduct,
			product,
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

		if err := scripts.EncodeJSON(w, product); err != nil {
			return
		}
	}
}

func ReadProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		products, err := products.ReadProducts(
			ctx,
			db,
			queries.ProductQueries.ReadProducts,
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

		if err := scripts.EncodeJSON(w, products); err != nil {
			return
		}
	}
}

func ReadProductsByOutlet(db *sql.DB) http.HandlerFunc {
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

		products, err := products.ReadProductsByOutlet(
			ctx,
			db,
			queries.ProductQueries.ReadProductsByOutlet,
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

		if err := scripts.EncodeJSON(w, products); err != nil {
			return
		}
	}
}

func ReadProductsByAvailable(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		available := r.PathValue("available")

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		products, err := products.ReadProductsByAvailable(
			ctx,
			db,
			queries.ProductQueries.ReadProductsByAvailable,
			available,
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

		if err := scripts.EncodeJSON(w, products); err != nil {
			return
		}
	}
}

func UpdateProductPrice(db *sql.DB) http.HandlerFunc {
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

		price, err := scripts.DecodeJSON[int64](
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

		err = products.UpdateProductPrice(
			ctx,
			db,
			queries.ProductQueries.UpdateProductPrice,
			price,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Product not found",
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

func UpdateProductAvailable(db *sql.DB) http.HandlerFunc {
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

		available, err := scripts.DecodeJSON[string](
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

		err = products.UpdateProductAvailable(
			ctx,
			db,
			queries.ProductQueries.UpdateProductAvailable,
			available,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Product not found",
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
