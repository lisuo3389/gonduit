package entities

// DiffusionCommit represents a commit in Diffusion.
type DiffusionCommit struct {
	ID             string `json:"id"`
	PHID           string `json:"phid"`
	RepositoryPHID string `json:"repositoryPHID"`
	Identifier     string `json:"identifier"`
	Epoch          string `json:"epoch"`
	URI            string `json:"uri"`
	IsImporting    bool   `json:"isImporting"`
	Summary        string `json:"summary"`
	AuthorPHID     string `json:"authorPHID"`
	CommitterPHID  string `json:"committerPHID"`
	Author         string `json:"author"`
	AuthorName     string `json:"authorName"`
	AuthorEmail    string `json:"authorEmail"`
	Committer      string `json:"committer"`
	CommitterName  string `json:"committerName"`
}

type RepositorySearchField struct {
	Name               string                 `json:"name"`
	Vcs                string                 `json:"vcs"`
	Callsign           string                 `json:"callsign"`
	ShortName          string                 `json:"shortName"`
	Status             string                 `json:"status"`
	IsImporting        bool                   `json:"isImporting"`
	AlmanacServicePHID string                 `json:"almanacServicePHID"`
	RefRules           map[string][]string    `json:"refRules"`
	DefaultBranch      string                 `json:"defaultBranch"`
	Description        string                 `json:"description"`
	SpacePHID          string                 `json:"spacePHID"`
	DateCreated        int64                  `json:"dateCreated"`
	DateModified       int64                  `json:"dateModified"`
	Policy             map[string]interface{} `json:"policy"`
}
