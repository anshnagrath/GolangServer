package controller

import (
	"encoding/json"
	"net/http"

	"../views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(data)
		}
	}
}
