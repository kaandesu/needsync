package needs

type Action struct {
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Title   string      `json:"title"`
	Options []string    `json:"options,omitempty"`
	Fields  []string    `json:"fields,omitempty"`
	Example string      `json:"example,omitempty"`
	Answer  interface{} `json:"answer"`
}

type IssueRef struct {
	Repo   string `json:"repo"`
	Number int    `json:"number"`
}

type Need struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Priority     string    `json:"priority"`
	Status       string    `json:"status"`
	Blocking     bool      `json:"blocking"`
	Description  string    `json:"description"`
	HumanActions []Action  `json:"human_actions"`
	Issue        *IssueRef `json:"issue,omitempty"`
}
