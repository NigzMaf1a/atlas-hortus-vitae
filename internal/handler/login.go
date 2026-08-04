package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/links"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/auth"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/outlets"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/queries"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode credentials
		creds, err := scripts.DecodeJSON[auth.HortusVirtaeCred](r.Body)

		if err != nil {
			fmt.Println("Error decoding login credentials:", err)

			http.Error(
				w,
				"Invalid login credentials",
				http.StatusBadRequest,
			)

			return
		}

		// Prepare request for Auth Service
		payload := auth.LoginCred{
			Email:    creds.Email,
			Password: creds.Password,
		}

		jsonData, err := json.Marshal(payload)

		if err != nil {
			fmt.Println("Error encoding login request:", err)

			http.Error(
				w,
				"Internal server error",
				http.StatusInternalServerError,
			)

			return
		}

		// Call Auth Service
		resp, err := http.Post(
			links.AuthLink,
			"application/json",
			bytes.NewBuffer(jsonData),
		)

		if err != nil {
			fmt.Println("Auth service request failed:", err)

			http.Error(
				w,
				"Authentication service unavailable",
				http.StatusBadGateway,
			)

			return
		}

		defer resp.Body.Close()

		// Auth service rejected credentials
		if resp.StatusCode != http.StatusOK {
			http.Error(
				w,
				"Invalid email or password",
				resp.StatusCode,
			)

			return
		}

		// Decode Auth Service response
		var loginResponse auth.LoginResponse

		if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
			fmt.Println("Error decoding auth service response:", err)

			http.Error(
				w,
				"Invalid authentication service response",
				http.StatusBadGateway,
			)

			return
		}

		// Create context for outlet lookup
		ctx, cancel := context.WithTimeout(
			r.Context(),
			5*time.Second,
		)

		defer cancel()

		// Get outlet
		outlet, err := outlets.ReadOutlet(
			ctx,
			db,
			queries.OutletQueries.ReadOutlet,
			creds.OutletID,
		)

		if err != nil {
			fmt.Println("Error reading outlet:", err)

			http.Error(
				w,
				"Failed to retrieve outlet",
				http.StatusInternalServerError,
			)

			return
		}

		// Build Hortus login response
		response := auth.HortusLoginResponse{
			Token:      loginResponse.Token,
			User:       loginResponse.User,
			OutletName: outlet.Name,
			OutletID:   outlet.OutletID,
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		if err := scripts.EncodeJSON(w, response); err != nil {
			fmt.Println("Error encoding login response:", err)

			http.Error(
				w,
				"Internal server error",
				http.StatusInternalServerError,
			)

			return
		}
	}
}
