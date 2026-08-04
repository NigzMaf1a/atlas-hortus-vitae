package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/payments"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/queries"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func CreatePayment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		payment, err := scripts.DecodeJSON[payments.Payment](r.Body)

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

		payment, err = payments.CreatePayment(
			ctx,
			db,
			queries.PaymentQueries.CreatePayment,
			payment,
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

		if err := scripts.EncodeJSON(w, payment); err != nil {
			return
		}
	}
}

func ReadPayments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)
		defer cancel()

		payments, err := payments.ReadPayments(
			ctx,
			db,
			queries.PaymentQueries.ReadPayments,
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

		if err := scripts.EncodeJSON(w, payments); err != nil {
			return
		}
	}
}

func ReadPaymentsByUser(db *sql.DB) http.HandlerFunc {
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

		payments, err := payments.ReadPaymentsByUser(
			ctx,
			db,
			queries.PaymentQueries.ReadPaymentsByUser,
			userID,
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

		if err := scripts.EncodeJSON(w, payments); err != nil {
			return
		}
	}
}

func UpdatePaymentStatus(db *sql.DB) http.HandlerFunc {
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

		err = payments.UpdatePaymentStatus(
			ctx,
			db,
			queries.PaymentQueries.UpdatePaymentStatus,
			status,
			id,
		)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(
					w,
					"Payment not found",
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