package main

import (
	"backend-assigment/common/app"
	"backend-assigment/common/postgresql"
	"backend-assigment/controller"
	"backend-assigment/persistence"
	"backend-assigment/service"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	// Create a new context
	ctx := context.Background()
	// Create a new Echo instance
	e := echo.New()

	// Initialize configuration manager
	configurationManager := app.NewConfigurationManager()

	// Create a PostgreSQL connection pool
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	// Initialize report repository
	reportRepository := persistence.NewReportRepository(dbPool)

	// Initialize report service
	reportService := service.NewReportService(reportRepository)

	// Initialize report controller
	reportController := controller.NewReportController(reportService)

	// Register routes for report controller
	reportController.RegisterRoutes(e)

	// Start the server
	err := e.Start(":8080")
	if err != nil {
		log.Errorf("Error while starting server: %v", err)
	}
}
