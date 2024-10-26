package utils

import (
	"log"
	"net/http"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[endpoint]: %s; [device]: %s, [method]: %v", r.URL.Path, r.UserAgent(), r.Method)
		f(w, r)
	}
}