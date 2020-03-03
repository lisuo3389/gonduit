package requests

type SearchMap map[string]interface{}

type SearchRequest struct {
	QueryKey    SearchMap `json:"query_key,omitempty"`
	Constraints SearchMap `json:"constraints,omitempty"`
	Attachments SearchMap `json:"attachments,omitempty"`
	Order       SearchMap `json:"order,omitempty"`
	Before      int       `json:"before,omitempty"`
	After       int       `json:"after,omitempty"`
	Limit       int       `json:"limit,omitempty"`
}
