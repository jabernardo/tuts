package handlers

import (
	"fmt"
	"net/http"
)

func ArticlesRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Reading article", r.PathValue("article"))
}
