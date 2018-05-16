package main

import (
	"fmt"
	"os"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	elogrus "gopkg.in/sohlich/elogrus.v3"
)

type ILogIO interface {
	Info()
	Error()
}
type LogIO struct{}

func (l LogIO) Info() {
	var log = logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.DebugLevel
	log.Out = os.Stdout

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		log.Panic(err)
	}

	hook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "golflog")

	if err != nil {
		log.Panic(err)
	}

	log.Hooks.Add(hook)

	// TODO: Refactor
	details := logrus.Fields{
		"stackTrace": logrus.Fields{
			"test":  true,
			"test2": logrus.Fields{},
		},
	}

	// TODO: Refactor
	log.WithFields(logrus.Fields{
		"serviceName": "go-validate",
		"statusCode":  200,
		"info":        details,
	}).Info("Hello world!")
}

func (l LogIO) Error() {
	fmt.Println(l)
}

// func logja() {
// 	var log = logrus.New()
// 	log.Formatter = &logrus.JSONFormatter{}
// 	log.Level = logrus.DebugLevel
// 	log.Out = os.Stdout

// 	client, err := elastic.NewClient(elastic.SetSniff(false))

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	hook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "golflog")

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	log.Hooks.Add(hook)

// 	// TODO: Refactor
// 	details := logrus.Fields{
// 		"stackTrace": logrus.Fields{
// 			"test":  true,
// 			"test2": logrus.Fields{},
// 		},
// 	}

// 	// TODO: Refactor
// 	log.WithFields(logrus.Fields{
// 		"serviceName": "go-validate",
// 		"statusCode":  200,
// 		"info":        details,
// 	}).Info("Hello world!")
// }

var details = logrus.Fields{
	"stackTrace": logrus.Fields{
		"test":  true,
		"test2": logrus.Fields{},
	},
}

// Interface type value
var logIO ILogIO = &LogIO{}
