package use

import (
	oa "github.com/apache/servicecomb-mesher/proxy/handler/oauth2"
	"golang.org/x/oauth2"
	"net/http"
)

var auth *oa.OAuth2

func Use(middleware *oa.OAuth2) {
	auth = middleware
}

func test(a *oa.OAuth2)  {
   a.GrantType = "Authorization_code"
   a.UseConfig = &oauth2.Config{
	   ClientID:     "",           // (required, string) your client_ID
	   ClientSecret: "",           // (required, string) your client_Secret
	   Scopes:       []string{""}, // (optional, string) scope specifies requested permissions
	   RedirectURL:  "",           // (required, string) URL to redirect users going through the OAuth2 flow
	   Endpoint: oauth2.Endpoint{ // (required, string) your auth server endpoint
		   AuthURL:  "",
		   TokenURL: "",
	   },
   }
   a.Authenticate = func(at string, req *http.Request) error {

		return nil
	}
}


