package main;;;:*+::;;;++***:::::;;:::**

import (
	"log"
	"net/http"
	"time"
)

type LogConfig struct {
	ServiceName   string
	StatusCode    string
	CorrelationID string
	Detail        logrus.Fields
}


func logger(inner http.Handler, name string) http.Handler {
	logConfig := LogConfig{
		ServiceName: "MYAPI",
		StatusCode: "200",
		CorrelationID: "xxxxx-xxxx-xxxx-xxx-xxx",
		Detail: {},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		// TODO: Feature toggle
		logIO.Info()

		log.Printf(
			"APP NAME# %s | 🚀 %s | 🔗 %s | 🏂 %s | ⏱️  %s",
			logConfig.ServiceName,
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
