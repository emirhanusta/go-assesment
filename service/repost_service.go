package service

import (
	"backend-assigment/dto/request"
	"backend-assigment/dto/response"
	"backend-assigment/persistence"
	"fmt"
	"strings"
)

// IReportService defines the interface for report service.
type IReportService interface {
	GetAllWithPagination(page int, size int, requestBody request.RequestBody) response.ResponseBody
}

// ReportService implements the IReportService interface.
type ReportService struct {
	reportRepository persistence.IReportRepository
}

// NewReportService creates a new instance of ReportService.
func NewReportService(reportRepository persistence.IReportRepository) IReportService {
	return &ReportService{reportRepository: reportRepository}
}

// GetAllWithPagination retrieves reports with pagination and filtering from the database.
func (reportService *ReportService) GetAllWithPagination(page int, size int, requestBody request.RequestBody) response.ResponseBody {
	// Redirecting query building steps to functions
	filterQuery := buildFilterQuery(requestBody.Filters)
	sortingQuery := buildSortingQuery(requestBody.Ordering)
	paginationQuery, page, size := buildPaginationQuery(page, size)

	// Combining all queries
	query := fmt.Sprintf("SELECT * FROM report_output %s %s %s", filterQuery, sortingQuery, paginationQuery)

	// Executing query in the database and returning the results
	reportOutputs := reportService.reportRepository.GetAllWithPagination(query)

	return response.ToResponse(reportOutputs, page, size)
}

// buildFilterQuery creates the filtering query.
func buildFilterQuery(filters []request.Filter) string {
	query := ""
	if len(filters) > 0 {
		query += " WHERE "
		for i, filter := range filters {
			if i != 0 {
				query += " AND "
			}
			// Determining filter type
			filterType := determineFilterType(filter.Value)
			// Building filter query
			query += buildSingleFilterQuery(filter, filterType)
		}
	}
	return query
}

// determineFilterType determines the type of filter.
func determineFilterType(value interface{}) string {
	switch value.(type) {
	case []interface{}:
		return "enum"
	case float64:
		return "number"
	case string:
		return "free_form"
	default:
		return ""
	}
}

// buildSingleFilterQuery creates a query for a single filter.
func buildSingleFilterQuery(filter request.Filter, filterType string) string {
	switch filterType {
	case "enum":
		// If the filter type is enum, create a query for an enum filter
		values := make([]string, len(filter.Value.([]interface{})))
		for i, v := range filter.Value.([]interface{}) {
			values[i] = v.(string)
		}
		// Construct the IN clause for enum values
		return fmt.Sprintf("%s IN ('%s')", filter.Column, strings.Join(values, "','"))
	case "number":
		// If the filter type is number, create a query for a number filter
		value := filter.Value.(float64)
		// Construct the query for number comparison
		return fmt.Sprintf("%s = %f", filter.Column, value)
	case "free_form":
		// If the filter type is free_form, create a query for a free-form filter
		value := filter.Value.(string)
		// Construct the LIKE clause for free-form matching
		return fmt.Sprintf("%s LIKE '%%%s%%'", filter.Column, value)
	default:
		// Return an empty string for unknown filter types
		return ""
	}
}

// buildSortingQuery creates the sorting query.
func buildSortingQuery(ordering []request.Ordering) string {
	// Initialize an empty query string
	query := ""
	// Check if there are any ordering criteria provided
	if len(ordering) > 0 {
		// If there are ordering criteria, add "ORDER BY" clause to the query
		query += " ORDER BY "
		// Iterate over each ordering criterion
		for i, ordering := range ordering {
			// If it's not the first criterion, add a comma to separate from the previous one
			if i != 0 {
				query += ", "
			}
			// Add the column name and direction to the query
			query += fmt.Sprintf("%s %s", ordering.Column, ordering.Direction)
		}
	}
	// Return the constructed query
	return query
}

// buildPaginationQuery creates the pagination query.
func buildPaginationQuery(page int, size int) (string, int, int) {
	// Determining page number and page size
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	// Building query
	return fmt.Sprintf(" LIMIT %d OFFSET %d", size, (page-1)*size), page, size
}
