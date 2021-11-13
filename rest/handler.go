package rest

import (
	"net/http"
	"skiresorts/business"
)

type handler struct {
	hs business.IHillsService
}

func New(hs business.IHillsService) *handler {
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
