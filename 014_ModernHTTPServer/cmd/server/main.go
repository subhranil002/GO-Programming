package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/subhranil002/GO-CRUD/internal/config"
	"github.com/subhranil002/GO-CRUD/internal/employee"
	"github.com/subhranil002/GO-CRUD/internal/database"
	"github.com/subhranil002/GO-CRUD/internal/middleware"
	"github.com/subhranil002/GO-CRUD/internal/router"
	"github.com/subhranil002/GO-CRUD/pkg/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.Load()
	if err != nil {
		log.Error("configuration error", "error", err)
		os.Exit(1)
	}

	startupCtx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancel()

	mongoClient, err := database.Connect(
		startupCtx,
		cfg.MongoURI,
	)
	if err != nil {
		log.Error("mongodb connection failed", "error", err)
		os.Exit(1)
	}

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := mongoClient.Disconnect(shutdownCtx); err != nil {
			log.Error(
				"mongodb disconnect failed",
				"error",
				err,
			)
		}
	}()

	collection := mongoClient.
		Database(cfg.MongoDatabase).
		Collection(cfg.MongoCollection)

	employeeRepo := course.NewRepository(collection)
	employeeService := course.NewService(employeeRepo)
	employeeHandler := course.NewHandler(employeeService)

	mux := router.Setup(employeeHandler)
	
	// Wrap the mux with the logger middleware
	loggedMux := middleware.RequestLogger(log)(mux)

	server := &http.Server{
		Addr:              ":" + strconv.Itoa(cfg.Port),
		Handler:           loggedMux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Info(
		"server started",
		"port",
		cfg.Port,
		"environment",
		cfg.AppEnv,
	)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error(
			"server stopped unexpectedly",
			"error",
			err,
		)
		os.Exit(1)
	}
}
