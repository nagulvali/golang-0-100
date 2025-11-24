package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK) // no need to set status as json already returns 200 OK

}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {

	var payload User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}

	// users = append(users, u)

	if err := insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Println(users)
	w.WriteHeader(http.StatusCreated)

}

func insertUser(u User) error {
	if u.FirstName == "" {
		return errors.New("FirstName is required")
	}

	if u.LastName == "" {
		return errors.New("LastName is required")
	}

	// storage validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("")
		}
	}

	users = append(users, u)
	return nil
}
