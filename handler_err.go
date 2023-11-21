package main

import (
	"net/http"

	"github.com/elsierra/go-server/util"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	util.RespondWithError(w, 400, "something went wrong")
}
