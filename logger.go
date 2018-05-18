package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type LogConfig struct {
	ServiceName   string
	RouteName     string
	RouteMethod   string
	RoutePattern  string
	StatusCode    string
	CorrelationID string
	Detail        log.Fields
}

// https://godoc.org/github.com/sirupsen/logrus
const Environment string = "dev"

func routeLogger(inner http.Handler, rtName string, rtMethod string, rtPattern string) http.Handler {
	if Environment == "production" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
	} else {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		routeLogging := LogConfig{
			ServiceName:   "MYAPI",
			RouteName:     rtName,
			RouteMethod:   rtMethod,
			RoutePattern:  rtPattern,
			StatusCode:    "200",
			CorrelationID: "xxxxx-xxxx-xxxx-xxx-xxx", // generating from api gateway or frontend
		}

		log.WithFields(
			log.Fields{
				"log":          routeLogging,
				"responseTime": time.Since(start),
			},
		).Info("Route serve")

	})
}

func logger(l interface{}) {
	fmt.Println(l)
}
