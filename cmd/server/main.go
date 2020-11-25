package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prima101112/crudsample/pkg/connection/db"
	"github.com/prima101112/crudsample/pkg/handler"
)

func main() {
	router := httprouter.New()
	db.InitDBConn()
	//ping
	router.GET("/ping", handler.Ping)

	//gopher
	router.GET("/gopher/:id", handler.JSON(handler.GetGopher))
	router.GET("/gophers", handler.JSON(handler.GetGophers))
	router.POST("/gopher", handler.JSON(handler.InsertGopher))
	router.PUT("/gopher/:id", handler.JSON(handler.UpdateGopher))

	//cors incase there is a frontend
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	//serve http
	log.Println("http.ListenAndServe in port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
