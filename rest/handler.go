package rest

import (
	"net/http"
	"skiresorts/usecase"
)

type handler struct {
	hs usecase.IHillsService
}

func New(hs usecase.IHillsService) *handler {
	return &handler{
		hs: hs,
	}
}

type Response struct {
	Status int
	Body   []byte
}

type IHandler interface {
	HandleHills(w http.ResponseWriter, r *http.Request)
}
