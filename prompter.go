package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type RepositoryResponse struct {
	User struct {
		Repositories struct {
			Nodes []struct {
				NameWithOwner string
				UpdatedAt     string
				Description   string
				Url           string
			}
		}
	}
}

type IssueResponse struct {
	User struct {
		Issues struct {
			Nodes []IssuePRNode
		}
	}
}

type PullRequestResponse struct {
	User struct {
		PullRequests struct {
			Nodes []IssuePRNode
		}
	}
}

type IssuePRNode struct {
	Title        string
	UpdatedAt    string
	ResourcePath string
	BodyText     string
	Url          string
}

func RepositoryPrompter(repo RepositoryResponse) {
	repoList := repo.User.Repositories.Nodes
	templates := &promptui.SelectTemplates{
		Label:    "  [ {{ . }} ]",
		Active:   "\U000026F3 {{ .NameWithOwner | red }}",
		Inactive: "  {{ .NameWithOwner | cyan }}",
		Selected: "\U000026F3 {{ .NameWithOwner | green | cyan }}",
		Details: `
--------- Repo ----------
{{ "Name:" | faint }}	{{ .NameWithOwner }}
{{ "Description:" | faint }}	{{ .Description }}
`,
	}

	prompt := promptui.Select{
		Label:     "Repositories",
		Items:     repoList,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", repoList[index].Url)
	OpenBrowser(repoList[index].Url)
}

func IssuePrompter(issue IssueResponse) {
	issueList := issue.User.Issues.Nodes

	templates := &promptui.SelectTemplates{
		Label:    "  [ {{ . }} ]",
		Active:   "\U0001F47F {{ .Title | red }} ({{ .UpdatedAt | green }})",
		Inactive: "  {{ .Title | cyan }} ({{ .UpdatedAt | green }})",
		Selected: "\U0001F47F {{ .Title | green | red }}",
		Details: `
--------- Repo ----------
{{ "Description:" | faint }}	{{ .BodyText }}
{{ "Resource:" | faint }}	{{ .ResourcePath }}
`,
	}

	reformedIssues := reformIssueData(issueList)

	prompt := promptui.Select{
		Label:     "Issues",
		Items:     reformedIssues,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Go to issue %q\n", reformedIssues[index].Title)
	OpenBrowser(reformedIssues[index].Url)
}

func PRPrompter(issue PullRequestResponse) {
	prList := issue.User.PullRequests.Nodes

	templates := &promptui.SelectTemplates{
		Label:    "  [ {{ . }} ]",
		Active:   "\U0001F47F {{ .Title | red }} ({{ .UpdatedAt | green }})",
		Inactive: "  {{ .Title | cyan }} ({{ .UpdatedAt | green }})",
		Selected: "\U0001F47F {{ .Title | green | red }}",
		Details: `
--------- Repo ----------
{{ "Description:" | faint }}	{{ .BodyText }}
{{ "Resource:" | faint }}	{{ .ResourcePath }}
`,
	}

	reformedPRs := reformIssueData(prList)

	prompt := promptui.Select{
		Label:     "Pull requests",
		Items:     reformedPRs,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Go to PR %q\n", reformedPRs[index].Title)
	OpenBrowser(reformedPRs[index].Url)
}

func reformIssueData(issueList []IssuePRNode) []IssuePRNode {
	var reformedNode []IssuePRNode
	for _, v := range issueList {
		GetPassedTime(v.UpdatedAt)
		reformedNode = append(reformedNode, IssuePRNode{
			Title:        v.Title,
			UpdatedAt:    GetPassedTime(v.UpdatedAt),
			ResourcePath: v.ResourcePath,
			BodyText:     TruncateLongText(v.BodyText),
			Url:          v.Url,
		})
	}

	return reformedNode
}
