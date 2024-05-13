package kernel

import (
	"context"
	"encoding/json"
	"eon/kata/mike/pkg/config"
	"eon/kata/mike/pkg/kernel/middleware"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	Server *http.Server
	Router *http.ServeMux
	Config *config.Config
	Logger *slog.Logger
}

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (app *Application) Run() {
	app.Logger.Info(fmt.Sprintf("Starting server on :%s", app.Config.App.Port))

	if err := app.Server.ListenAndServe(); err != nil {
		app.Logger.Error(err.Error())
		panic(err)
	}
}

func Boot() *Application {
	conf := config.Load()

	router := http.NewServeMux()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	application := &Application{
		Router: router,
		Server: &http.Server{
			Addr:         ":" + conf.App.Port,
			Handler:      middleware.WithLogging(router, logger),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Config: conf,
		Logger: logger,
	}

	return application
}

func (app *Application) Respond(response http.ResponseWriter, request *http.Request, data interface{}, status int) {
	response.Header().Set("Content-Type", app.Config.HTTP.Content)

	wrappedData := ResponseData{
		Status: status,
		Data:   data,
	}

	response.WriteHeader(status)

	err := json.NewEncoder(response).Encode(wrappedData)
	if err != nil {
		app.Logger.Error(err.Error())
		panic(err)
	}
}

// WaitForShutdown - Lets wait for a shutdown signal and shutdown gracefully
func (app *Application) WaitForShutdown() {
	// Create a channel to listen for OS signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal through our channel
	<-interruptChan

	app.Logger.Info("Received shutdown signal, gracefully terminating")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.Server.Shutdown(ctx)
	os.Exit(1)
}
