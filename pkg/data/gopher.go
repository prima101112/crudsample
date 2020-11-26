package data

import (
	"strings"

	"github.com/google/uuid"
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
	var g entity.Gophers
	err := db.GetCon().Select(&g, "SELECT * FROM gophers")
	if err != nil {
		return g, err
	}
	return g, nil
}

func InsertGophers(g entity.Gopher) (string, error) {
	//need pagination limit offset here
	//whatever id is on request or not it will replace with uuid
	id := generateID()
	_, err := db.GetCon().Exec("INSERT INTO gophers (id, name, email, company) VALUES ($1, $2, $3, $4)", id, g.Name, g.Email, g.Company)
	if err != nil {
		return "", err
	}
	return id, nil
}

func UpdateGophers(id string, gnew entity.Gopher) (entity.Gopher, error) {
	var g entity.Gopher
	gold, err := GetGopher(id)
	if err != nil {
		return g, err
	}
	if gnew.Name == "" {
		gnew.Name = gold.Name
	}
	if gnew.Email == "" {
		gnew.Email = gold.Email
	}
	if gnew.Company == "" {
		gnew.Company = gold.Company
	}

	_, err = db.GetCon().Exec("UPDATE gophers SET name = $2, email = $3, company = $4 WHERE id=$1;", id, gnew.Name, gnew.Email, gnew.Company)
	if err != nil {
		return gold, err
	}
	return gnew, nil
}

func generateID() string {
	u := uuid.New()
	uuid := strings.Replace(u.String(), "-", "", -1)
	return uuid
}
