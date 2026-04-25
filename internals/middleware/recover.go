package middleware

import (
	"fmt"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic: ", r)

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Something went wrong on our end."))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
