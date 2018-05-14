package main

import (
	"log"
	"net/http"
	"time"
	// "os"
	// elastic "github.com/olivere/elastic"
	// "github.com/sirupsen/logrus"
	// "gopkg.in/sohlich/elogrus.v3"
)

func main() {
	router := InitRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

	// var log = logrus.New()
	// log.Formatter = &logrus.JSONFormatter{}
	// log.Level = logrus.DebugLevel
	// log.Out = os.Stdout

	// client, err := elastic.NewClient(elastic.SetSniff(false))
	// if err != nil {
	// 	log.Panic(err)
	// }
	// hook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "mylog")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Hooks.Add(hook)

	// details := logrus.Fields{
	// 	"stackTrace": logrus.Fields{
	// 		"test":  true,
	// 		"test2": logrus.Fields{},
	// 	},
	// }
	// log.WithFields(logrus.Fields{
	// 	"serviceName": "go-validate",
	// 	"statusCode":  200,
	// 	"info":        details,
	// }).Info("Hello world!")
}
