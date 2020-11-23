package asana

import (
	"encoding/json"

	"github.com/fsmiamoto/asana/domain"
)

type TasksResponse struct {
	Data []domain.Task `json:"data"`
}

func (c *Client) GetTasksFromProject(projectGID string) ([]domain.Task, error) {
	req, err := c.buildRequest("GET", "/projects/"+projectGID+"/tasks", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body TasksResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Data, nil
}
