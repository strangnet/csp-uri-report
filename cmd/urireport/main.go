package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/strangnet/csp-uri-report/internal/pkg/http"
	"github.com/strangnet/csp-uri-report/internal/pkg/inmem"
	"github.com/strangnet/csp-uri-report/internal/pkg/report"
)

const (
	defaultAPIAddress = ":8080"
)

func main() {

	var (
		apiAddress = envString("API_ADDRESS", defaultAPIAddress)
	)

	flag.Parse()

	logger := log.StandardLogger()
	logger.SetFormatter(&log.JSONFormatter{})

	errorChannel := make(chan error)

	var reports = inmem.NewReportRepository()
	var rs = report.NewService(reports)

	go func() {
		log.WithField("addr", apiAddress).Info("http.server.listen")

		server := http.NewServer(logger, apiAddress, rs)

		errorChannel <- server.Open()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got signal: %s", <-c)
	}()

	if err := <-errorChannel; err != nil {
		log.Error(errors.New("got error: " + err.Error()))
	}

	log.Info("terminated")
}

func envString(key, defaultValue string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		return defaultValue
	}
	return value
}
