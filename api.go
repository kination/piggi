package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

func GetRepositories() (RepositoryResponse, error) {
	client := defaultGraphQLConnection()
	repoRequest := graphql.NewRequest(`
        query getRepos($login: String!, $last: Int!){
            user(login: $login) {
                repositories(last: $last orderBy: {
                    field:UPDATED_AT
                    direction: ASC
                    }) {
                    nodes {
                        nameWithOwner
                        description
                        url
                    }
                }
            }
        }
    `)

	SetupRequest(repoRequest)
	repoRequest.Var("last", 10)

	ctx := context.Background()

	var response RepositoryResponse
	if err := client.Run(ctx, repoRequest, &response); err != nil {
		log.Fatal(err)
		return response, err
	}

	return response, nil
}

func SetupRequest(req *graphql.Request) {
	user, err := ConfigReader()
	if err != nil {
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
