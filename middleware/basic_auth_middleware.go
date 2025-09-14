package middleware

import (
	"log"
	"net/http"
)

type BasicAuthMiddleware struct {
	Next http.Handler
}

func (bam *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if bam.Next == nil {
		bam.Next = http.DefaultServeMux
	}

	if r.Method != http.MethodGet {
		username, password, ok := r.BasicAuth()
		if !ok {
			log.Println("Basic auth parsing error")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != "test" {
			log.Println("Invalid username")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if password != "123" {
			log.Println("Invalid password")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	bam.Next.ServeHTTP(w, r)

}
