package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/labstack/echo-contrib/echoprometheus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Info struct {
	Hostname string `json:"hostname" xml:"hostname"`
	Commit   string `json:"commit" xml:"commit"`
	Info     string `json:"additionalInfo" xml:"additionalInfo"`
}

type AppVersion struct {
  Version string `json:"version" xml:"version"`
}

type AppStatus struct {
  Status string `json:"status" xml:"status"`
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000"}))

	e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
		Subsystem: "otusgoapp",
		LabelFuncs: map[string]echoprometheus.LabelValueFunc{
			"app": func(c echo.Context, err error) string { // additional custom label
				return "otusgoapp"
			},
		},
	}))

	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/metrics", echoprometheus.NewHandler())
	e.GET("/version", version)
	e.GET("/ready", ready)
	e.GET("/alive", alive)

	e.Logger.SetLevel(log.INFO)
	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func hello(c echo.Context) error {

	hostname := os.Getenv("HOSTNAME")
	commit := os.Getenv("COMMIT")
	additionalInfo := os.Getenv("ADDITIONAL_INFO")

	if len(hostname) == 0 {
		commit = "Ooops! We dont have commit sha!"
	}

	if len(commit) == 0 {
		commit = "Ooops! We dont have commit sha!"
	}

	if len(additionalInfo) == 0 {
		additionalInfo = "Ooops! We dont additional info message!"
	}

	answer := &Info{
		Hostname: hostname,
		Commit:   commit,
		Info:     additionalInfo,
	}

	return c.JSONPretty(http.StatusOK, answer, "  ")

}

func version(c echo.Context) error {

  version:= os.Getenv("VERSION")
  if len(version) == 0 {
    version = "Ooops! We dont have version!"
  }

  answer := &AppVersion{
		Version: version,
	}

  return c.JSONPretty(http.StatusOK, answer, "  ")
}

func ready (c echo.Context) error {

  status:= "OK"
  answer := &AppStatus{
    Status: status,
  }

  return c.JSONPretty(http.StatusOK, answer, "  ")
}

func alive (c echo.Context) error {

  status:= "OK"
  answer := &AppStatus{
    Status: status,
  }

  return c.JSONPretty(http.StatusOK, answer, "  ")
}