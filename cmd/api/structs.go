package main

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrResp struct {
	Error string
}
