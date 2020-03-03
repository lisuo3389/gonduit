package responses

type Cursor struct {
	Limit int `json:"limit"`
	After *int `json:"after"`
	Before *int `json:"before"`
	Order map[string]string `json:"order"`
}
