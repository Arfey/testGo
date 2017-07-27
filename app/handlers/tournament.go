package handlers

import (
	"net/http"
	"encoding/json"
	"../models"
	"../consts"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		fmt.Println(a, e)
		if a == e {
			return true
		}
	}
	return false
}

// Tournament's handlers

func CreateTournament(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	var tournament models.Tournament
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tournament)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	defer r.Body.Close()
	err = models.CreateTournament(tournament.Deposit)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Reset(w http.ResponseWriter, r *http.Request) {
	models.Reset()
}

func SetWinner(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	var winner models.Winner
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&winner)
	vars := mux.Vars(r)
	tournaments_id, _ := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	defer r.Body.Close()
	err = models.SetWinner(winner.Id, tournaments_id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func AddMember(w http.ResponseWriter, r *http.Request) {
	initJson(w)
	var member models.Member
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&member)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	fmt.Println(member, err)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})

		return
	}

	iscontains := contains(
		member.Backers,
		member.User_id,
	)

	if iscontains {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf(consts.UserBackerInclude)
		return
	}

	deposit, err := models.GetTournamentDeposit(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})

		return
	}

	min, count, err := models.GetMembersData(
		member.User_id,
		member.Backers,
	)

	if len(member.Backers) + 1 != count {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{consts.UserNotFound})

		return
	}

	prise := int(deposit / count)

	fmt.Println(prise)

	if min < prise {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{consts.LessMinDepositLimit})

		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.
			NewEncoder(w).
			Encode(Error{err.Error()})

		return
	}

	models.SetUserPrise(member.User_id, member.Backers, prise)
	models.AddMember(member.User_id , id, member.Backers, prise)

	w.WriteHeader(http.StatusCreated)
}