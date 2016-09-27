package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/newrelic/go-agent"
)

var monitoringApp newrelic.Application

func monitorApplication() newrelic.Application {
	if monitoringApp != nil {
		return monitoringApp
	}
	config := monitorConfig()
	var err error
	monitoringApp, err = newrelic.NewApplication(config)
	if err != nil {
		log.Fatal(err)
	}
	return monitoringApp
}

func monitorConfig() newrelic.Config {
	appName := GetEnv("APP_NAME", true)
	licKey := GetEnv("NEW_RELIC_LICENSE_KEY", true)

	config := newrelic.NewConfig(appName, licKey)
	config.Enabled = configEnabled()
	return config
}

func configEnabled() bool {
	disabled := strings.ToLower(strings.TrimSpace(GetEnv("MONITORING_DISABLED", false)))
	if disabled == "true" {
		log.Println("Monitoring disabled for this environment.")
		return false
	}
	return true
}

func monitorHandle(pattern string, handler http.Handler) (string, http.Handler) {
	app := monitorApplication()
	return newrelic.WrapHandle(app, pattern, handler)
}

func MonitorMiddleware(pattern string, fn func(http.ResponseWriter, *http.Request), o ServerOptions) (string, http.Handler) {
	return monitorHandle(pattern, Middleware(fn, o))
}

func MonitorImageMiddleware(o ServerOptions) func(string, Operation) (string, http.Handler) {
	return func(pattern string, fn Operation) (string, http.Handler) {
		return monitorHandle(pattern, validateImage(Middleware(imageController(o, Operation(fn)), o), o))
	}
}
