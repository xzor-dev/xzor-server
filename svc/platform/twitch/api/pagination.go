package api

// Pagination holds information on pagination from API queries.
type Pagination struct {
	Cursor string `json:"cursor"`
}
