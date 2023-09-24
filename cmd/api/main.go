package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/user/insert", insertUser)

	err := http.ListenAndServe(os.Getenv("HTTP_PORT"), nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resp := ErrResp{
			Error: "Method not allowed",
		}
		respond(w, resp, http.StatusMethodNotAllowed)
		return
	}

	req, err := decodeRequest(r)
	if err != nil {
		respond(
			w,
			ErrResp{
				Error: err.Error(),
			},
			http.StatusBadRequest,
		)
		return
	}

	isValid, why := isValidRequest(req)
	if !isValid {
		respond(
			w,
			ErrResp{
				Error: why,
			},
			http.StatusBadRequest,
		)
		return
	}

	user, err := saveUser(req)
	if err != nil {
		respond(
			w,
			ErrResp{
				Error: err.Error(),
			},
			http.StatusInternalServerError,
		)
		return
	}

	respond(w, user, http.StatusCreated)
}

func decodeRequest(r *http.Request) (Request, error) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func respond(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
