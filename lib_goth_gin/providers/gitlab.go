package providers

import "github.com/markbates/goth/providers/gitlab"

// https://gitlab.com/oauth/applications/ to get clientId...

func NewGitlabProvider() *gitlab.Provider {
	clientId := "09cb4ef656412ea42a934e1f5582836dd18119711221990024ba80ff352d210e"
	clientSecrets := "1682c88cacc66cee7e3b8b7ca5fe2cfde1c880156383f8b2ff519bac0ffcea70"
	authCallbackUrl := "http://localhost:9090/auth/callback/gitlab"
	return gitlab.New(clientId, clientSecrets, authCallbackUrl)
}
