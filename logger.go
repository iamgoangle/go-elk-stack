package main

import (
	"log"
	"net/http"
	"time"
)

type LogConfig struct {
	App string
}

func logger(inner http.Handler, name string) http.Handler {
	logConfig := LogConfig{
		App: "GOLF ggez",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		logja()

		log.Printf(
			"APP NAME# %s | 🚀 %s | 🔗 %s | 🏂 %s | ⏱️  %s",
			logConfig.App,
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
