package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/xzor-dev/xzor-server/svc/team"
)

func main() {
	router := initRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         ":1337",
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("server has shut down: %v", err)
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()
	initTeamRouter(router)
	return router
}

func initTeamRouter(r *mux.Router) {
	router := r.PathPrefix("/teams").Subrouter()
	service := team.NewService()
	team.RegisterHTTPRoutes(service, router)
}
