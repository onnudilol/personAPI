package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host   = "localhost"
	port   = 32768
	user   = "postgres"
	dbname = "postgres"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host string, port int, user string, dbname string) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	var err error

	a.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = a.DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/api/teams/{name}", a.getTeams).Methods("GET")
	a.Router.HandleFunc("/api/members/{name}", a.getMembers).Methods("GET")
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
}

func (a *App) getTeams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	teams, err := getTeams(a.DB, name)

	if err != nil {
		log.Println(err)
	}

	JSONResponse(w, http.StatusOK, teams)
}

func (a *App) getMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	members, err := getMembers(a.DB, name)

	if err != nil {
		log.Println(err)
	}

	JSONResponse(w, http.StatusOK, members)
}

func JSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
