package controllers

import (
	"fmt"
	"net/http"
)

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All my events working with air")
}
