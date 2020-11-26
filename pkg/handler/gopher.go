package handler

import (
	"encoding/json"
	"log"
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
		log.Printf("%s", err)
		res.Status = "failed"
		res.Message = "error get gopher"
		json.NewEncoder(w).Encode(res)
		return
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
		log.Printf("%s", err)
		res.Status = "failed"
		res.Message = "error get gopher"
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Status = "success"
	res.Message = "success get gophers"
	json.NewEncoder(w).Encode(res)
}

func InsertGopher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var res entity.DefaultResponse
	decoder := json.NewDecoder(r.Body)
	var g entity.Gopher
	err := decoder.Decode(&g)
	if err != nil {
		log.Printf("%s", err)
		res.Status = "failed"
		res.Message = "failed to decode request body"
		json.NewEncoder(w).Encode(res)
		return
	}

	id, err := data.InsertGophers(g)
	g.ID = id
	if err != nil {
		log.Printf("%s", err)
		res.Status = "failed"
		res.Message = "error insert gopher"
		json.NewEncoder(w).Encode(res)
	}
	res.Data = g
	res.Status = "success"
	res.Message = "success inert gophers"
	json.NewEncoder(w).Encode(res)
}

func UpdateGopher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var res entity.DefaultResponse
	decoder := json.NewDecoder(r.Body)
	var g entity.Gopher
	err := decoder.Decode(&g)
	if err != nil {
		log.Printf("%s", err)
		res.Status = "failed"
		res.Message = "failed to decode request body"
		json.NewEncoder(w).Encode(res)
		return
	}
	gnew, err := data.UpdateGophers(id, g)
	gnew.ID = id
	if err != nil {
		log.Printf("%s", err)
		res.Status = "failed"
		res.Message = "error update gopher"
		json.NewEncoder(w).Encode(res)
	}
	res.Data = gnew
	res.Status = "success"
	res.Message = "success update gophers"
	json.NewEncoder(w).Encode(res)
}
