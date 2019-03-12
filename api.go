package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/machinebox/graphql"
)

func ghCall() {
	gqClient := graphql.NewClient("https://api.github.com/graphql")
	gqClient.Log = func(s string) { log.Println(s) }

	repoRequest := graphql.NewRequest(`query getRepos($login: String!, $first: Int!){
		user(login: $login) {
		  repositories(first: $first) {
			nodes {
			  name
			  description
			}
		  }
		}
	  }`)
	repoRequest.Var("login", "...")
	repoRequest.Var("first", 10)

	user, err := configReader()
	if err != nil {
		return
	}

	repoRequest.Header.Set("Authorization", "Bearer "+user.Token)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := gqClient.Run(ctx, repoRequest, &respData); err != nil {
		log.Fatal(err)
		return
	}

	prettier, err := json.MarshalIndent(respData, "", "  ")
	if err != nil {
		return
	}

	log.Println("Response: " + string(prettier))
}
