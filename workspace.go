package asana

import (
	"encoding/json"

	"github.com/fsmiamoto/asana/domain"
)

type WorkspacesResponse struct {
	Data []domain.Workspace `json:"data"`
}

type WorkspaceResponse struct {
	Data domain.Workspace `json:"data"`
}

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

	var body WorkspacesResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Data, nil
}

func (c *Client) GetWorkspace(id string) (*domain.Workspace, error) {
	req, err := c.buildRequest("GET", "/workspaces/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body WorkspaceResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return &body.Data, nil
}
