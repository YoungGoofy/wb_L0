package route

import (
	"context"
	"encoding/json"
	"github.com/YoungGoofy/WB_L0/internal/services/db"
	"github.com/gorilla/mux"
	"html/template"
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

func (h *Handler) OpenAPI(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api" {
		http.NotFound(w, r)
		return
	}
	htmlPath := "/home/frosty/go/src/github.com/YoungGoofy/WB_L0/internal/services/server/ui/html/index.html"
	ts, err := template.ParseFiles(htmlPath)
	if err != nil {
		h.log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusBadGateway)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		h.log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusBadGateway)
	}
	err = r.ParseForm()
	if err != nil {
		return
	}
	uid := r.FormValue("inputOrderUID")
	if uid == "" {
		h.log.Print("no id")
		http.Error(w, "No UID", http.StatusBadRequest)
	} else {
		order, err := h.mix.GetById(context.Background(), uid)
		if err != nil {
			h.log.Print("no id")
			http.Error(w, "No order with this uid", http.StatusBadRequest)
		} else {
			err = json.NewEncoder(w).Encode(order)
			if err != nil {
				return
			}
		}
	}
}
