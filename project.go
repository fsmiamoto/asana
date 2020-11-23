package asana

import (
	"encoding/json"

	"github.com/fsmiamoto/asana/domain"
)

type ProjectsResponse struct {
	Data []domain.Project `json:"data"`
}

func (c *Client) GetWorkspaceProjects(workspaceID string) ([]domain.Project, error) {
	req, err := c.buildRequest("GET", "/workspaces/"+workspaceID+"/projects", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body ProjectsResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Data, nil
}
