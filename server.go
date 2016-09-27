package main

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type ServerOptions struct {
	Port             int
	Burst            int
	Concurrency      int
	HttpCacheTtl     int
	HttpReadTimeout  int
	HttpWriteTimeout int
	CORS             bool
	Gzip             bool
	EnableURLSource  bool
	AuthForwarding   bool
	Address          string
	ApiKey           string
	Mount            string
	CertFile         string
	KeyFile          string
	Authorization    string
	AlloweOrigins    []*url.URL
}

func Server(o ServerOptions) error {
	addr := o.Address + ":" + strconv.Itoa(o.Port)
	handler := NewLog(NewServerMux(o), os.Stdout)

	server := &http.Server{
		Addr:           addr,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Duration(o.HttpReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(o.HttpWriteTimeout) * time.Second,
	}

	return listenAndServe(server, o)
}

func listenAndServe(s *http.Server, o ServerOptions) error {
	if o.CertFile != "" && o.KeyFile != "" {
		return s.ListenAndServeTLS(o.CertFile, o.KeyFile)
	}
	return s.ListenAndServe()
}

func NewServerMux(o ServerOptions) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(MonitorMiddleware("/", indexController, o))
	mux.Handle(MonitorMiddleware("/form", formController, o))
	mux.Handle(MonitorMiddleware("/health", healthController, o))

	image := MonitorImageMiddleware(o)
	mux.Handle(image("/resize", Resize))
	mux.Handle(image("/enlarge", Enlarge))
	mux.Handle(image("/extract", Extract))
	mux.Handle(image("/crop", Crop))
	mux.Handle(image("/rotate", Rotate))
	mux.Handle(image("/flip", Flip))
	mux.Handle(image("/flop", Flop))
	mux.Handle(image("/thumbnail", Thumbnail))
	mux.Handle(image("/zoom", Zoom))
	mux.Handle(image("/convert", Convert))
	mux.Handle(image("/watermark", Watermark))
	mux.Handle(image("/info", Info))

	return mux
}

// func monitorApplication() newrelic.Application {
// 	config := newrelic.NewConfig("gsc-svc-image-proc-dev", "27df0df4ce963c493f2ff0550ca4fffde361e873")
// 	app, err := newrelic.NewApplication(config)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	return app
// }
