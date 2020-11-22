package asana

import (
	"encoding/json"

	"github.com/fsmiamoto/asana/domain"
)

func (c *Client) GetWorkspaces() ([]domain.Workspace, error) {
	req, err := c.buildRequest("GET", "/workspaces", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body domain.WorkspaceResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Data, nil
}
