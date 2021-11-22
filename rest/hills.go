package rest

import (
	"encoding/json"
	"net/http"
	"skiresorts/domain"
	"strconv"
	"strings"
)

type hillDTO struct {
	Length float64 `json:"length"`
	Slope  float64 `json:"slope"`
	Height float64 `json:"height"`
}

func toHillDTO(model *domain.Hill, height float64) *hillDTO {
	return &hillDTO{
		Length: model.Length,
		Slope:  model.Slope,
		Height: height,
	}
}

func getHillID(path string) int {
	idString := strings.TrimPrefix(path, "/hills/")
	id, _ := strconv.Atoi(idString)
	return id
}

func (h *handler) HandleHills(w http.ResponseWriter, r *http.Request) {
	id := getHillID(r.URL.Path)
	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		resp := h.getHill(id)
		w.WriteHeader(resp.Status)
		w.Write(resp.Body)

	case "POST":

	case "PATCH":
		var hill domain.Hill
		err := json.NewDecoder(r.Body).Decode(&hill)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		resp := h.updateHill(id, &hill)
		w.WriteHeader(resp.Status)
		w.Write(resp.Body)
	}
}

func (h *handler) updateHill(id int, hill *domain.Hill) *Response {
	err := h.hs.UpdateHill(id, hill)
	if err != nil {
		return &Response{Status: http.StatusBadRequest}
	}
	return &Response{Status: http.StatusOK}
}

func (h *handler) getHill(id int) *Response {
	hill, err := h.hs.GetHill(id)
	var encoded []byte
	var errObj struct {
		E string `json:"error"`
	}
	if err == nil {
		encoded, err = json.Marshal(toHillDTO(hill, hill.GetHeight()))
		return &Response{
			Status: http.StatusOK,
			Body:   encoded,
		}
	}
	errObj.E = err.Error()
	encoded, err = json.Marshal(errObj)
	return &Response{
		Status: http.StatusBadRequest,
		Body:   encoded,
	}
}
