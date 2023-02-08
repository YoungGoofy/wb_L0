package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) Route() {
	r := mux.NewRouter()
	r.HandleFunc("/api", h.OpenAPI)
	r.HandleFunc("/api/{id}", h.GetIUD).Methods("GET")
	h.log.Fatal(http.ListenAndServe("localhost:8080", r))
}
