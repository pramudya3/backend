package cmd

import (
	"fmt"

	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func InitSupertokens(cfg Config) {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// https://try.supertokens.com is for demo purposes. Replace this with the address of your core instance (sign up on supertokens.com), or self host a core.
			ConnectionURI: fmt.Sprintf("%v", cfg.SupertokensURI),
			APIKey:        "someKey",
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "backend",
			APIDomain:       "http://localhost:3000",
			WebsiteDomain:   "http://localhost:1234",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			session.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
