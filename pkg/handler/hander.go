package handler

import (
	"fmt"
	"net/http"
)

func TopPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
