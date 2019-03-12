package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/machinebox/graphql"
)

func Repositories() {
	client := defaultGraphQLConnection()
	repoRequest := graphql.NewRequest(`
        query getRepos($login: String!, $last: Int!){
            user(login: $login) {
            repositories(last: $last) {
                    nodes {
                        nameWithOwner
                        description
                    }
                }
            }
        }
    `)

	SetupRequest(repoRequest)
	repoRequest.Var("last", 5)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, repoRequest, &respData); err != nil {
		log.Fatal(err)
		return
	}

	prettier, err := json.MarshalIndent(respData, "", "  ")
	if err != nil {
		return
	}

	log.Println("Response: " + string(prettier))
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
	graphQLClient.Log = func(s string) { log.Println(s) }

	return graphQLClient
}
