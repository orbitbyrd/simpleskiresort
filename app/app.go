package app

import (
	"net/http"
	"regexp"
	"skiresorts/business"
	"skiresorts/persistence"
	"skiresorts/rest"
)

type Mode int

const (
	InMem        Mode = 0
	FromPostgres Mode = 1
)

func Run(mode Mode) {
	var repo persistence.IHillRepo
	switch mode {
	case InMem:
		repo = persistence.NewInMemRepo()
	case FromPostgres:
		repo = persistence.NewPostgresRepo()
	}
	hs := business.NewHillService(repo)

	restHandler := rest.New(hs)

	/*hill, err := hs.GetHill(1)
	if err == nil {

		fmt.Printf("Hills height = %f\n", hill.GetHeight())
	}*/

	regexHandler := &RegexpHandler{}
	regexHandler.HandleFunc(regexp.MustCompile(`^/hills/\d+$`), restHandler.HandleHills)
	http.ListenAndServe(":9999", regexHandler)
}
