package needs

import (
	"fmt"
	"needsync/internal/git"
	"path/filepath"
)

func SyncCreate(dir string) error {
	client := git.New()

	needs, err := LoadAll(dir)
	if err != nil {
		return err
	}
	fmt.Printf("nice %+v\n", needs)

	for _, n := range needs {

		fmt.Println(n)

		// skip if already linked
		if n.Issue != nil && n.Issue.Number != 0 {
			continue
		}

		body := buildIssueBody(n)
		fmt.Printf("> %+v\n", body)

		issue, err := client.CreateIssue(n.Title, body)
		if err != nil {
			return err
		}
		fmt.Printf(">issue no: %d\n", issue.Number)

		n.Issue = &IssueRef{
			Repo:   client.Owner + "/" + client.Repo,
			Number: issue.Number,
		}

		n.Status = "waiting_human"

		path := filepath.Join(dir, n.ID+".json")
		if err := Save(path, n); err != nil {
			return err
		}

		fmt.Println("Created issue for:", n.ID)
	}

	return nil
}

func buildIssueBody(n *Need) string {
	body := fmt.Sprintf("## %s\n\n%s\n\n", n.Title, n.Description)

	for _, a := range n.HumanActions {
		body += fmt.Sprintf("### %s\n", a.Title)

		switch a.Type {
		case "decision":
			body += "Options:\n"
			for _, o := range a.Options {
				body += "- " + o + "\n"
			}
		case "input":
			body += "Fields:\n"
			for _, f := range a.Fields {
				body += "- " + f + "\n"
			}
		case "confirmation":
			body += "Example: " + a.Example + "\n"
		}

		body += "\n"
	}

	return body
}
