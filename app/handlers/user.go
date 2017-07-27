package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"../models"
)

// Router handlers for User model

func GetUser(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}
	user, err := models.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.
		NewEncoder(w).
		Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})

		return
	}

	err = models.DeleteUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	var user models.User
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	models.UpdateUser(id, user.Balance)

	if err != nil {
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})

		return
	}

	w.WriteHeader(200)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	err := models.DeleteUsers()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	defer r.Body.Close()
	err = models.CreateUser(user.Balance)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
