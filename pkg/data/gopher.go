package data

import (
	"github.com/prima101112/crudsample/pkg/connection/db"
	"github.com/prima101112/crudsample/pkg/entity"
)

//never use * is real life. this is just example

func GetGopher(id string) (entity.Gopher, error) {
	g := entity.Gopher{}
	err := db.GetCon().Get(&g, "SELECT * FROM gophers WHERE ID=$1", id)
	if err != nil {
		return g, err
	}
	return g, nil
}

func GetGophers() (entity.Gophers, error) {
	//need pagination limit offset here
	g := []entity.Gopher{}
	err := db.GetCon().Get(&g, "SELECT * FROM gophers")
	if err != nil {
		return g, err
	}
	return g, nil
}
