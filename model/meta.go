package model

import "time"

type AuthorizedUser struct {
	User AuthorizedUserDetail `json:"user"`
}

type AuthorizedUserDetail struct {
	ID        string `json:"_id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AuthorizationInfo struct {
	ID          string        `json:"_id"`
	CreatedOn   time.Time     `json:"createdOn"`
	GrantType   string        `json:"grantType"`
	LastUsed    time.Time     `json:"lastUsed"`
	Sites       []interface{} `json:"sites"`
	Orgs        []string      `json:"orgs"`
	Workspaces  []interface{} `json:"workspaces"`
	Users       []string      `json:"users"`
	RateLimit   int           `json:"rateLimit"`
	Status      string        `json:"status"`
	Application Application   `json:"application,omitempty"`
}

type Application struct {
	ID          string `json:"_id"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	OwnerType   string `json:"ownerType"`
}
