package providers

import "github.com/markbates/goth"

func Init() {
	goth.UseProviders(NewGithubProvider())
	goth.UseProviders(NewGitlabProvider())
}
