package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

func GetRepositories() (RepositoryResponse, error) {
	client := defaultGraphQLConnection()
	repoRequest := graphql.NewRequest(`
        query getRepos($login: String!, $first: Int!){
            user(login: $login) {
                repositories(first: $first orderBy: {
                    field:UPDATED_AT
                    direction: DESC
                    }) {
                    nodes {
						nameWithOwner
						updatedAt
                        description
                        url
                    }
                }
            }
        }
    `)

	setupRequest(repoRequest)
	repoRequest.Var("first", 10)

	ctx := context.Background()

	var response RepositoryResponse
	if err := client.Run(ctx, repoRequest, &response); err != nil {
		log.Fatal(err)
		return response, err
	}

	return response, nil
}

func GetIssues() (IssueResponse, error) {
	client := defaultGraphQLConnection()
	issueRequest := graphql.NewRequest(`
        query getIssues($login: String!, $first: Int!){
            user(login: $login) {
                issues(first: $first states: OPEN orderBy:{
                    field:UPDATED_AT
                    direction: DESC
                }) {
                    nodes {
						title
						updatedAt
						resourcePath
						bodyText
                        url
                    }
                }
            }
        }
    `)

	setupRequest(issueRequest)
	issueRequest.Var("first", 20)

	ctx := context.Background()

	var response IssueResponse
	if err := client.Run(ctx, issueRequest, &response); err != nil {
		log.Fatal(err)
		return response, err
	}

	return response, nil
}

func GetPullRequests() (PullRequestResponse, error) {
	client := defaultGraphQLConnection()
	prRequest := graphql.NewRequest(`
        query getIssues($login: String!, $first: Int!){
            user(login: $login) {
                pullRequests(first: $first states: OPEN orderBy:{
                    field:UPDATED_AT
                    direction: DESC
                }) {
                    nodes {
						title
						updatedAt
						resourcePath
						bodyText
                        url
                    }
                }
            }
        }
    `)

	setupRequest(prRequest)
	prRequest.Var("first", 20)

	ctx := context.Background()

	var response PullRequestResponse
	if err := client.Run(ctx, prRequest, &response); err != nil {
		log.Fatal(err)
		return response, err
	}

	return response, nil
}

func GetUserInfo() (UserResponse, error) {
	client := defaultGraphQLConnection()
	userRequest := graphql.NewRequest(`
        query getUser($login: String!){
            user(login: $login) {
				bio
				email
				createdAt
				company
            }
        }
    `)

	setupRequest(userRequest)

	ctx := context.Background()

	var response UserResponse
	if err := client.Run(ctx, userRequest, &response); err != nil {
		log.Fatal(err)
		return response, err
	}

	return response, nil
}

func setupRequest(req *graphql.Request) {
	user, err := ConfigReader()
	if err != nil {
		// TODO: error handler
		return
	}

	req.Var("login", user.Id)
	req.Header.Set("Authorization", "Bearer "+user.Token)
}

func defaultGraphQLConnection() *graphql.Client {
	graphQLClient := graphql.NewClient("https://api.github.com/graphql")
	// graphQLClient.Log = func(s string) { log.Println(s) }

	return graphQLClient
}
