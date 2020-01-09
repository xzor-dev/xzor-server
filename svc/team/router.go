package team

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xzor-dev/xzor-server/lib/api"
)

type router struct {
	service *Service
}

// RegisterHTTPRoutes adds all available team routes to the router.
func RegisterHTTPRoutes(s *Service, r *mux.Router) {
	tr := &router{
		service: s,
	}
	r.HandleFunc("", api.JSONResponder(tr.getTeams)).Methods("GET")
	r.HandleFunc("", api.JSONResponder(tr.createTeam)).Methods("POST")
	r.HandleFunc("/{teamID}", api.JSONResponder(tr.getTeam)).Methods("GET")
}

func (r *router) createTeam(res http.ResponseWriter, req *http.Request) *api.JSONResponse {
	params := &CreateTeamParams{}
	err := json.NewDecoder(req.Body).Decode(params)
	if err != nil {
		return api.NewJSONErrorResponse(err)
	}

	if data := params.Validate(); data != nil {
		return api.NewJSONFailResponse(nil, data)
	}

	team, err := r.service.CreateTeam(params)
	if err != nil {
		return api.NewJSONErrorResponse(err)
	}
	return api.NewJSONResponse(team)
}

func (r *router) getTeam(res http.ResponseWriter, req *http.Request) *api.JSONResponse {
	vars := mux.Vars(req)
	teamID := ID(vars["teamID"])
	team, err := r.service.Team(teamID)
	if err != nil {
		return api.NewJSONErrorResponse(err)
	}
	return api.NewJSONResponse(team)
}

func (r *router) getTeams(res http.ResponseWriter, req *http.Request) *api.JSONResponse {
	teams, err := r.service.Teams()
	if err != nil {
		return api.NewJSONErrorResponse(err)
	}
	return api.NewJSONResponse(teams)
}
