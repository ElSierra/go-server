package main

import (
	"net/http"

	"github.com/elsierra/go-server/util"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w,200,struct{}{})
}
