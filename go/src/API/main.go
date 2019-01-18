package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	adsHandler "API/pkg/handlers/ads"
)

var PORT = ":8000"

func main() {
	router := mux.NewRouter()

	adsEndpoint := router.PathPrefix(`/{ads:ads\/?}`).Subrouter()
	// adsEndpoint.HandleFunc("", adsHandler.Find).Methods("GET")
	adsEndpoint.HandleFunc("", adsHandler.Create).Methods("POST")
	// adsEndpoint.HandleFunc("/{id}", adsHandler.FindByID).Methods("GET")
	adsEndpoint.HandleFunc("/{id}", adsHandler.Edit).Methods("PUT")
	adsEndpoint.HandleFunc("/{id}", adsHandler.Delete).Methods("DELETE")
	adsEndpoint.HandleFunc("/{id}/ad-params", adsHandler.Edit).Methods("PUT")

	fmt.Println("Running on http://localhost" + PORT)
	http.ListenAndServe(PORT, responseJson(router))
}

func responseJson(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
}