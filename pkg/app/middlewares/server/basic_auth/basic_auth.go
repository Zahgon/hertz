package basic_auth

import (
	"github.com/cloudwego/hertz/pkg/app"
)

type Accounts map[string]string

type pairs map[string]string

func (p pairs) findValue(needle string) (v string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func constructPairs(accounts Accounts) pairs { _ = "STUB: not implemented"; return *new(pairs) }

func BasicAuthForRealm(accounts Accounts, realm, userKey string) app.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(app.HandlerFunc)
}

func BasicAuth(accounts Accounts) app.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(app.HandlerFunc)
}
