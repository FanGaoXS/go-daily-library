package providers

import "github.com/markbates/goth/providers/gitlab"

// https://gitlab.com/oauth/applications/ to get clientId...

func NewGitlabProvider() *gitlab.Provider {
	clientId := "fc2d410be64f6af61b47fb81b40c0e7ddaee9f227f9ebbcea36363bd1c1583eb"
	clientSecrets := "040a92904ad45511b380d23b7df973dbc52cc4eef3cee947a9ac9d0da46f1ac5"
	authCallbackUrl := "http://localhost:9090/auth/callback/gitlab"
	return gitlab.New(clientId, clientSecrets, authCallbackUrl)
}
