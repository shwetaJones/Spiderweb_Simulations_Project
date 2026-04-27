package serpentapi

// SearchResult represents a single organic search result from the Serpent API.
type SearchResult struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

// SearchResponse represents the top-level JSON response from the Serpent API.
type SearchResponse struct {
	SearchParameters SearchParameters `json:"search_parameters"`
	OrganicResults   []SearchResult   `json:"organic_results"`
	Error            string           `json:"error,omitempty"`
}

// SearchParameters echoes back the query parameters used in the request.
type SearchParameters struct {
	Engine string `json:"engine"`
	Query  string `json:"q"`
}
