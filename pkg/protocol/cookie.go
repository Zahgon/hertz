package protocol

import (
	"sync"
	"time"

	"github.com/cloudwego/hertz/internal/nocopy"
	"github.com/cloudwego/hertz/pkg/common/errors"
)

const (
	CookieSameSiteDisabled CookieSameSite = iota

	CookieSameSiteDefaultMode

	CookieSameSiteLaxMode

	CookieSameSiteStrictMode

	CookieSameSiteNoneMode
)

var zeroTime time.Time

var (
	errNoCookies = errors.NewPublic("no cookies found")

	CookieExpireDelete = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	CookieExpireUnlimited = zeroTime
)

type CookieSameSite int

type Cookie struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	key    []byte
	value  []byte
	expire time.Time
	maxAge int
	domain []byte
	path   []byte

	httpOnly bool
	secure   bool

	partitioned bool
	sameSite    CookieSameSite

	bufKV argsKV
	buf   []byte
}

var cookiePool = &sync.Pool{
	New: func() interface{} {
		return &Cookie{}
	},
}

func AcquireCookie() *Cookie { _ = "STUB: not implemented"; return nil }

func ReleaseCookie(c *Cookie) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetDomain(domain string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetPath(path string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetPathBytes(path []byte) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetExpire(expire time.Time) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetKey(key string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetKeyBytes(key []byte) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetValue(value string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetValueBytes(value []byte) { _ = "STUB: not implemented"; return }

func (c *Cookie) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendCookiePart(dst, key, value []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendRequestCookieBytes(dst []byte, cookies []argsKV) []byte {
	_ = "STUB: not implemented"
	return nil
}

func appendResponseCookieBytes(dst []byte, cookies []argsKV) []byte {
	_ = "STUB: not implemented"
	return nil
}

type cookieScanner struct {
	b []byte
}

func parseRequestCookies(cookies []argsKV, src []byte) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func (s *cookieScanner) next(kv *argsKV) bool { _ = "STUB: not implemented"; return false }

func (c *Cookie) Key() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) Cookie() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) Reset() { _ = "STUB: not implemented"; return }

func (c *Cookie) Value() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) Parse(src string) error { _ = "STUB: not implemented"; return nil }

func (c *Cookie) ParseBytes(src []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Cookie) MaxAge() int { _ = "STUB: not implemented"; return 0 }

func (c *Cookie) SetMaxAge(seconds int) { _ = "STUB: not implemented"; return }

func (c *Cookie) Expire() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *Cookie) Domain() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) Path() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) Secure() bool { _ = "STUB: not implemented"; return false }

func (c *Cookie) SetSecure(secure bool) { _ = "STUB: not implemented"; return }

func (c *Cookie) SameSite() CookieSameSite { _ = "STUB: not implemented"; return *new(CookieSameSite) }

func (c *Cookie) Partitioned() bool { _ = "STUB: not implemented"; return false }

func (c *Cookie) SetSameSite(mode CookieSameSite) { _ = "STUB: not implemented"; return }

func (c *Cookie) HTTPOnly() bool { _ = "STUB: not implemented"; return false }

func (c *Cookie) SetHTTPOnly(httpOnly bool) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetPartitioned(partitioned bool) { _ = "STUB: not implemented"; return }

func (c *Cookie) String() string { _ = "STUB: not implemented"; return "" }

func decodeCookieArg(dst, src []byte, skipQuotes bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func getCookieKey(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func warnIfInvalid(value []byte) bool { _ = "STUB: not implemented"; return false }
