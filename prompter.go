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

func RepositoryPrompter(repo RepositoryResponse) {
	repoList := repo.User.Repositories.Nodes
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F645 {{ .NameWithOwner | cyan }}",
		Inactive: "  {{ .NameWithOwner | cyan }}",
		Selected: "\U0001F645 {{ .NameWithOwner | green | cyan }}",
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
