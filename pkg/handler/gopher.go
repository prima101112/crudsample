package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prima101112/crudsample/pkg/data"
	"github.com/prima101112/crudsample/pkg/entity"
)

func GetGopher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var res entity.DefaultResponse
	gop, err := data.GetGopher(id)
	res.Data = gop

	if err != nil {
		res.Status = "failed"
		res.Message = "error get gopher"
	}
	res.Status = "success"
	res.Message = "success get gopher"
	json.NewEncoder(w).Encode(res)
}
func GetGophers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var res entity.DefaultResponse
	gop, err := data.GetGophers()
	res.Data = gop

	if err != nil {
		res.Status = "failed"
		res.Message = "error get gopher"
	}
	res.Status = "success"
	res.Message = "success get gophers"
	json.NewEncoder(w).Encode(res)
}
func InsertGopher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var res entity.DefaultResponse
	gop, err := data.GetGophers()
	res.Data = gop

	if err != nil {
		res.Status = "failed"
		res.Message = "error get gopher"
	}
	res.Status = "success"
	res.Message = "success get gophers"
	json.NewEncoder(w).Encode(res)
}
func UpdateGopher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var res entity.DefaultResponse
	gop, err := data.GetGophers()
	res.Data = gop

	if err != nil {
		res.Status = "failed"
		res.Message = "error get gopher"
	}
	res.Status = "success"
	res.Message = "success get gophers"
	json.NewEncoder(w).Encode(res)
}
