package responses

import "github.com/uber/gonduit/entities"

type DiffusionRepositorySearchResponse struct {
	Data   []*DiffusionRepositorySearchData `json:"data"`
	Maps   map[string]interface{}        `json:"maps"`
	Query  map[string]interface{}        `json:"query"`
	Cursor *Cursor                       `json:"cursor"`
}

type DiffusionRepositorySearchData struct {
	ID     string                            `json:"id"`
	Type   string                            `json:"type"`
	PHID   string                            `json:"phid"`
	Fields *entities.RepositorySearchField `json:"fields"`
}
