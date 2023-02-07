package route

import (
	"context"
	"encoding/json"
	"github.com/YoungGoofy/WB_L0/internal/services/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	mix db.Repositories
	log *log.Logger
}

func NewHandler(mix db.Repositories, logger *log.Logger) *Handler {
	return &Handler{mix: mix, log: logger}
}

func (h *Handler) GetIUD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if params["id"] == "" {
		h.log.Print("no id")
		http.Error(w, "No UID", http.StatusBadRequest)
	} else {
		order, err := h.mix.GetById(context.Background(), params["id"])
		//fmt.Println(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			err = json.NewEncoder(w).Encode(order)
			if err != nil {
				return
			}
		}
	}
}
