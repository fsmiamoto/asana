package domain

type Workspace struct {
	GID            string    `json:"gid"`
	Name           string    `json:"name"`
	ResourceType   string    `json:"resource_type"`
	EmailDomains   *[]string `json:"email_domains"`
	IsOrganization *bool     `json:"is_organization"`
}
