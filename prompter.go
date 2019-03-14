package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type RepositoryResponse struct {
	User struct {
		Repositories struct {
			Nodes []struct {
				Description   string
				NameWithOwner string
				Url           string
			}
		}
	}
}

type IssueResponse struct {
	User struct {
		Issues struct {
			Nodes []IssueNode
		}
	}
}

type IssueNode struct {
	Title        string
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
		Active:   "\U0001F47F {{ .Title | red }}",
		Inactive: "  {{ .Title | cyan }}",
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

	fmt.Printf("Go to issue %q\n", issueList[index].Title)
	// OpenBrowser(issueList[index].Url)
}

func reformIssueData(issueList []IssueNode) []IssueNode {
	var reformedNode []IssueNode
	for _, v := range issueList {
		reformedNode = append(reformedNode, IssueNode{
			Title:        v.Title,
			ResourcePath: v.ResourcePath,
			BodyText:     TruncateString(v.BodyText),
			Url:          v.Url,
		})
	}

	return reformedNode
}
