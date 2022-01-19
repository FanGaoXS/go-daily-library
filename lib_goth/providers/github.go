package providers

import "github.com/markbates/goth/providers/github"

const (
	ClientId        = "7cf93272a9f64f52e0fa"
	ClientSecrets   = "cf62bab0300dee7c133041e3e366b11d3963cea4"
	AuthCallbackUrl = "http://localhost:9090/auth/github/callback"
)

func NewGithubProvider() *github.Provider {
	return github.New(ClientId, ClientSecrets, AuthCallbackUrl)
}
