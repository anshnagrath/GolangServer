package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../model"
	"../views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error occured"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {

			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				fmt.Println(err)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		} else if r.Method == http.MethodDelete {
			name := strings.Split(r.URL.Path, "/")[2]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"Item deleted"})

		}
	}
}
