package main

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

func RepoFormatter(repo RepositoryResponse) {
	// TODO
}
