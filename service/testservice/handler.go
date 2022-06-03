package testservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AllGroceries(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var req Request
	json.Unmarshal(reqBody, &req)

	h.service.AllGroceries(w, req)

}
