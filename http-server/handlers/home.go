package handlers

import (
	"fmt"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
