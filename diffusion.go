package gonduit

import (
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
)

// DiffusionQueryCommits performs a call to diffusion.querycommits.
func (c *Conn) DiffusionQueryCommits(
	req requests.DiffusionQueryCommitsRequest,
) (*responses.DiffusionQueryCommitsResponse, error) {
	var res responses.DiffusionQueryCommitsResponse

	if err := c.Call("diffusion.querycommits", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DiffusionRepositorySearch performs a call to diffusion.repository.search.
func (c *Conn)DiffusionRepositorySearch(
	req requests.DiffusionRepositorySearchRequest,
	) (*responses.DiffusionRepositorySearchResponse, error) {
	var res responses.DiffusionRepositorySearchResponse

	if err := c.Call("diffusion.repository.search", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
