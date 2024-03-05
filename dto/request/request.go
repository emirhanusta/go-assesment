package request

// Filter represents a filtering criterion.
type Filter struct {
	Column string      `json:"column"` // Column name to filter
	Type   string      `json:"type"`   // Type of filter: enum, number, free_form
	Value  interface{} `json:"value"`  // Value to filter by; interface{} is used for variable data type
}

// Ordering represents an ordering criterion.
type Ordering struct {
	Column    string `json:"column"`    // Column name to order by
	Direction string `json:"direction"` // Direction of ordering: ASC or DESC
}

// RequestBody represents the structure of request body.
type RequestBody struct {
	Filters  []Filter   `json:"filters"`  // Array of filters
	Ordering []Ordering `json:"ordering"` // Array of ordering criteria
}
