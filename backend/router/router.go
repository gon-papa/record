package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gon-papa/record/util/logger"
	"github.com/gon-papa/record/util/response"
)

func NewMux() http.Handler {
	r := chi.NewRouter()
	r.Use(accessLogger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		res := response.NewResponse("OK", nil)
		res.CreateResponse(w, 200)

	})
	return r
}

func accessLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		remoteAddr := r.RemoteAddr
		// リバースプロキシなどからアクセスしている場合は元のIPを取得する
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			remoteAddr = forwardedFor
		}
		logger.AccessLog("ip", remoteAddr, "method", r.Method, "path", r.URL.Path, "proto", r.Proto, "host", r.Host, "time", time.Since(start))
	})
}
