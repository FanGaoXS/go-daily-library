package providers

import "github.com/markbates/goth/providers/github"

// https://github.com/settings/developers to get clientId...

func NewGithubProvider() *github.Provider {
	clientId := "7cf93272a9f64f52e0fa"
	clientSecrets := "cf62bab0300dee7c133041e3e366b11d3963cea4"
	authCallbackUrl := "http://localhost:9090/auth/callback/github"
	return github.New(clientId, clientSecrets, authCallbackUrl)
}
