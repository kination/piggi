package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	piggi := cli.NewApp()
	piggi.Action = terminalAction

	err := piggi.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func terminalAction(c *cli.Context) error {
	subject := c.Args().Get(0)

	switch subject {
	case "repo":
		printoutSubject("Repositories")
		repo, err := GetRepositories()
		if err != nil {
			return err
		}

		RepositoryPrompter(repo)

	case "issue":
		printoutSubject("Issues")
		issue, err := GetIssues()
		if err != nil {
			return err
		}

		IssuePrompter(issue)

	case "pr":
		printoutSubject("Pull Requests")
		pr, err := GetPullRequests()
		if err != nil {
			panic(err)
		}

		PRPrompter(pr)

	// TODO: Prompter for Notification
	/*
		case "noti":
			printoutSubject("Notifications")
			log.Println("selected noti")
			// defaultAPIConnection()
	*/
	case "user":
		printoutUser()

	default:
		log.Println("Wrong selection!")
	}

	return nil
}

func printoutSubject(subject string) {
	fmt.Println("\n==================================")
	fmt.Println("\n Look over " + subject + "")
	fmt.Println("\n==================================\n")
}
