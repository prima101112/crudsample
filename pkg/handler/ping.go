package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "pong\n version 0.0.9")
}
