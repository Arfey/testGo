package handlers

import (
	"net/http"
	"../consts"
)

type Error consts.Error

func initJson(w http.ResponseWriter)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
