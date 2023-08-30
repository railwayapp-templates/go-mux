package health

import "net/http"

func Handler(w http.ResponseWriter, _ *http.Request) {
	// send a simple 200 status code back, this is where you could ping your database, etc
	w.WriteHeader(200)
}
