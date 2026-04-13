package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Client struct {
	BaseURL string
	Token   string
	Owner   string
	Repo    string
}

func New() *Client {
	return &Client{
		BaseURL: os.Getenv("GITEA_BASE_URL"),
		Token:   os.Getenv("GITEA_TOKEN"),
		Owner:   os.Getenv("GITEA_OWNER"),
		Repo:    os.Getenv("GITEA_REPO"),
	}
}

type Issue struct {
	Number int `json:"number"`
}

func (c *Client) CreateIssue(title, body string) (*Issue, error) {
	url := fmt.Sprintf("%s/api/v1/repos/%s/%s/issues", c.BaseURL, c.Owner, c.Repo)

	payload := map[string]interface{}{
		"title": title,
		"body":  body,
	}

	b, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue Issue
	json.NewDecoder(resp.Body).Decode(&issue)

	return &issue, nil
}
