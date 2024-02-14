package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const WEBFINGER_ACCOUNT_PREFIX string = "acct:"
const WEBFINGER_RESPONSE_TEMPLATE string = `{
  "subject": "acct:${email}",
  "links": [
    {
      "rel": "http://openid.net/specs/connect/1.0/issuer",
      "href": "${issuer_url}"
    }
  ]
}`

type serveWebFinger struct{}

func (s serveWebFinger) handle(w http.ResponseWriter, r *http.Request) {
	// resourceParam must be an URL encoded email address prefixed with "acct:".
	resourceParam := r.URL.Query().Get("resource")
	if resourceParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	webfingerQuery, err := url.QueryUnescape(resourceParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(webfingerQuery, WEBFINGER_ACCOUNT_PREFIX) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	account := strings.Split(webfingerQuery, WEBFINGER_ACCOUNT_PREFIX)[1]

	accountParts := strings.Split(account, "@")
	if len(accountParts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	domain := accountParts[1]
	if !isValidDomain(domain) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, buildResponse(account, "https://auth.thisismy.cloud"))
}

func isValidDomain(domain string) bool {
	switch domain {
	case
		"codehat.de",
		"thisismy.cloud",
		"skobigs.de":
		return true
	}
	return false
}

func buildResponse(email string, issuerUrl string) string {
	return strings.Replace(strings.Replace(WEBFINGER_RESPONSE_TEMPLATE, "${email}", email, 1), "${issuer_url}", issuerUrl, 1)
}
