package server

import (
	"net/http"

	"github.com/vfg2006/oauth-go/authenticator"
)

type Request struct {
	Text string `json:"text"`
}

func RequestToken(authenticatorService authenticator.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error in request", http.StatusBadRequest)
			return
		}

		text, err := authenticatorService.GetToken(req.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(text); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
