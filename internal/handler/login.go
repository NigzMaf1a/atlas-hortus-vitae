package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/links"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/auth"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/scripts"
)

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		creds, err := scripts.DecodeJSON[auth.HortusVirtaeCred](r.Body)

		if err != nil {
			fmt.Println("An error occurred while decoding login credentials")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"email":    creds.Email,
			"password": creds.Password,
		}

		jsonData, err := json.Marshal(data)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := http.Post(
			links.AuthLink,
			"application/json",
			bytes.NewBuffer(jsonData),
		)

		if err != nil {
			fmt.Println("Request failed:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := scripts.EncodeJSON(w, body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
