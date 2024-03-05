package controller

import (
	"backend-assigment/dto/request"
	"backend-assigment/dto/response"
	"backend-assigment/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// ReportController handles HTTP requests related to reports.
type ReportController struct {
	reportService service.IReportService
}

// NewReportController creates a new instance of ReportController.
func NewReportController(reportService service.IReportService) *ReportController {
	return &ReportController{reportService: reportService}
}

// RegisterRoutes registers routes for ReportController.
func (reportController *ReportController) RegisterRoutes(e *echo.Echo) {
	e.POST("/assignment/query", reportController.GetAllWithPagination)
}

// GetAllWithPagination handles requests to get reports with pagination.
func (reportController *ReportController) GetAllWithPagination(c echo.Context) error {
	// Parsing request body
	requestBody := new(request.RequestBody)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.ErrorResponse{Message: "Invalid request body"})
	}

	// Parsing page number
	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	if pageErr != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.ErrorResponse{Message: "Invalid page number"})
	}

	// Parsing page size
	size, sizeErr := strconv.Atoi(c.QueryParam("size"))
	if sizeErr != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.ErrorResponse{Message: "Invalid page size"})
	}

	// Getting reports with pagination from service
	responseBody := reportController.reportService.GetAllWithPagination(page, size, *requestBody)
	return c.JSON(http.StatusOK, responseBody)
}
