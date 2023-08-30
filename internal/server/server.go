package server

import (
	"app/internal/router"
	"net/http"
	"time"
)

func StartServer(address string) error {
	srv := &http.Server{
		Handler: router.MuxRouter,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		// adjust as needed
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return srv.ListenAndServe()
}
