package routes

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	v1 := http.NewServeMux()

	v1.Handle("/events/", RegisterEventsRoutes("/events"))

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))
}
