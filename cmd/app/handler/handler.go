package handler

import (
	"cafe_role/internal/controller"
	"cafe_role/internal/controller/req"
	"cafe_role/internal/controller/res"
	"cafe_role/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	m := mux.NewRouter()
	h := Handler{c: c}
	m.HandleFunc("/roles/{cafeId:[0-9]+}", h.getList).Methods(http.MethodGet)
	m.HandleFunc("/roles/{cafeId:[0-9]+}", h.create).Methods(http.MethodPost)
	m.HandleFunc("/roles/{cafeId:[0-9]+}/{roleId:[0-9]+}", h.patch).Methods(http.MethodPatch)
	m.HandleFunc("/roles/{cafeId:[0-9]+}/{roleId:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return m
}

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	reqPage := page.GetPageReqByRequest(r)
	list, total, err := h.c.GetList(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listTotalDto := res.ToListTotalDto(list, total)
	data, err := json.Marshal(listTotalDto)
	if err != nil {
		log.Println("getList json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	var d req.CreateDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Create(r.Context(), cafeId, d)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	roleId, err := strconv.Atoi(vars["roleId"])
	if err != nil {
		http.Error(w, "invalid role id", http.StatusBadRequest)
		return
	}
	var d req.PatchDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Patch(r.Context(), cafeId, roleId, d)
	if err != nil {
		if strings.Contains(err.Error(), "no row") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}

	roleId, err := strconv.Atoi(vars["roleId"])
	if err != nil {
		http.Error(w, "invalid role id", http.StatusBadRequest)
		return
	}

	err = h.c.Delete(r.Context(), cafeId, roleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
