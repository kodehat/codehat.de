package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/kodehat/codehat.de/pkg/mailcow"
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

	mailcowHost, hasMailcowHost := os.LookupEnv("MAILCOW_HOST")
	mailcowApiKey, hasMailcowApiKey := os.LookupEnv("MAILCOW_API_KEY")

	if hasMailcowHost && hasMailcowApiKey {
		mailcowClient := mailcow.MailcowClient{
			Host:   mailcowHost,
			ApiKey: mailcowApiKey,
		}
		isValidMailbox, err := isValidMailbox(mailcowClient, account)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		if !isValidMailbox {
			w.WriteHeader(http.StatusForbidden)
			return
		}
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

var mailboxesCache = ttlcache.New[string, bool](
	ttlcache.WithTTL[string, bool](time.Hour * 1),
)

func init() {
	go mailboxesCache.Start()
}

func isValidMailbox(client mailcow.MailcowClient, mailbox string) (bool, error) {
	validMailboxInCache := mailboxesCache.Has(mailbox)
	if validMailboxInCache {
		return mailboxesCache.Get(mailbox).Value(), nil
	}
	mailboxData, err := client.GetMailbox(mailbox)
	if err != nil {
		return false, err
	}
	validMailboxInCache = mailboxData.Username == mailbox
	mailboxesCache.Set(mailbox, validMailboxInCache, ttlcache.DefaultTTL)
	return (*mailboxData).Username == mailbox, nil
}

func buildResponse(email string, issuerUrl string) string {
	return strings.Replace(strings.Replace(WEBFINGER_RESPONSE_TEMPLATE, "${email}", email, 1), "${issuer_url}", issuerUrl, 1)
}
