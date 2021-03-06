package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func JSON(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		h(w, r, ps)
	}
}
