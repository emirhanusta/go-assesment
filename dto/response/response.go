package response

import "backend-assigment/domain"

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Message string `json:"message"` // Message describing the error
}

// ResponseBody represents the structure of response body.
type ResponseBody struct {
	Page     int           `json:"page"`      // Current page number
	PageSize int           `json:"page_size"` // Number of items per page
	Count    int           `json:"count"`     // Total count of results
	Results  []interface{} `json:"results"`   // Results array
}

// ToResponse converts a slice of report outputs to a ResponseBody.
func ToResponse(reports []domain.ReportOutput, page int, size int) ResponseBody {
	responseBody := ResponseBody{
		Page:     page,
		PageSize: size,
		Count:    len(reports),
		Results:  make([]interface{}, len(reports)),
	}
	for i, report := range reports {
		responseBody.Results[i] = report
	}
	return responseBody
}
