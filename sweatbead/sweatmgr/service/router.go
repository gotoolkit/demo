package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
	appName       = "sweatmgr"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps Dependencies) (router *mux.Router) {

	router = mux.NewRouter()

	// No version requirement for /ping
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	router.HandleFunc("/users", getUsersHandler).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/user", createUserHandler).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/user/sweat", getSweatByUserIdHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/sweat", createSweatHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/sweat_samples", getSweatSamplesHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)

	// Version 2 API management
	v2 := fmt.Sprintf("application/vnd.%s.v2", appName)

	router.HandleFunc("/user/sweat", getSweatByUserIdHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v2)
	router.HandleFunc("/sweat", createSweatHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v2)
	router.HandleFunc("/sweat_samples", getSweatSamplesHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v2)

	return
}
