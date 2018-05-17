package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	ServiceName   string
	RouteName     string
	RouteMethod   string
	RoutePattern  string
	StatusCode    string
	CorrelationID string
	Detail        logrus.Fields
}

func routeLogger(inner http.Handler, rtName string, rtMethod string, rtPattern string) http.Handler {
	logConfig := LogConfig{
		ServiceName:   "MYAPI",
		RouteName:     rtName,
		RouteMethod:   rtMethod,
		RoutePattern:  rtPattern,
		StatusCode:    "200",
		CorrelationID: "xxxxx-xxxx-xxxx-xxx-xxx", // generating from api gateway or frontend
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		log.Printf(
			"APP NAME# %s | 🚀 %s | 🔗 %s | 🏂 %s | ⏱️  %s",
			logConfig.ServiceName,
			r.Method,
			r.RequestURI,
			rtName,
			time.Since(start),
		)
	})
}

func logger(l interface{}) {
	fmt.Println(l)
}
