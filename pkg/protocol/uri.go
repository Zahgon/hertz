package protocol

import (
	"sync"

	"github.com/cloudwego/hertz/internal/nocopy"
)

func AcquireURI() *URI { _ = "STUB: not implemented"; return nil }

func ReleaseURI(u *URI) { _ = "STUB: not implemented"; return }

var uriPool = &sync.Pool{
	New: func() interface{} {
		return &URI{}
	},
}

type URI struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	pathOriginal []byte
	scheme       []byte
	path         []byte
	queryString  []byte
	hash         []byte
	host         []byte

	queryArgs       Args
	parsedQueryArgs bool

	DisablePathNormalizing bool

	fullURI    []byte
	requestURI []byte

	username []byte
	password []byte
}

type argsKV struct {
	key     []byte
	value   []byte
	noValue bool
}

func (kv *argsKV) GetKey() []byte { _ = "STUB: not implemented"; return nil }

func (kv *argsKV) GetValue() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) CopyTo(dst *URI) { _ = "STUB: not implemented"; return }

func (u *URI) QueryArgs() *Args { _ = "STUB: not implemented"; return nil }

func (u *URI) parseQueryArgs() { _ = "STUB: not implemented"; return }

func (u *URI) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetHash(hash string) { _ = "STUB: not implemented"; return }

func (u *URI) SetHashBytes(hash []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Username() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetUsername(username string) { _ = "STUB: not implemented"; return }

func (u *URI) SetUsernameBytes(username []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Password() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetPassword(password string) { _ = "STUB: not implemented"; return }

func (u *URI) SetPasswordBytes(password []byte) { _ = "STUB: not implemented"; return }

func (u *URI) QueryString() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetQueryString(queryString string) { _ = "STUB: not implemented"; return }

func (u *URI) SetQueryStringBytes(queryString []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Path() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetPath(path string) { _ = "STUB: not implemented"; return }

func (u *URI) String() string { _ = "STUB: not implemented"; return "" }

func (u *URI) SetPathBytes(path []byte) { _ = "STUB: not implemented"; return }

func (u *URI) PathOriginal() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) Scheme() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetScheme(scheme string) { _ = "STUB: not implemented"; return }

func (u *URI) SetSchemeBytes(scheme []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Reset() { _ = "STUB: not implemented"; return }

func (u *URI) Host() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetHost(host string) { _ = "STUB: not implemented"; return }

func (u *URI) SetHostBytes(host []byte) { _ = "STUB: not implemented"; return }

func (u *URI) LastPathSegment() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) Update(newURI string) { _ = "STUB: not implemented"; return }

func (u *URI) UpdateBytes(newURI []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Parse(host, uri []byte) { _ = "STUB: not implemented"; return }

func getScheme(rawURL []byte) (scheme, path []byte) { _ = "STUB: not implemented"; return nil, nil }

func (u *URI) parse(host, uri []byte, isTLS bool) { _ = "STUB: not implemented"; return }

func stringContainsCTLByte(s []byte) bool { _ = "STUB: not implemented"; return false }

func splitHostURI(host, uri []byte) ([]byte, []byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func normalizePath(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func copyArgs(dst, src []argsKV) []argsKV { _ = "STUB: not implemented"; return nil }

func (u *URI) updateBytes(newURI, buf []byte) []byte { _ = "STUB: not implemented"; return nil }

func isAbsoluteURI(uri []byte) bool { _ = "STUB: not implemented"; return false }

func (u *URI) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) RequestURI() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) appendSchemeHost(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) FullURI() []byte { _ = "STUB: not implemented"; return nil }

func ParseURI(uriStr string) *URI { _ = "STUB: not implemented"; return nil }

type Proxy func(*Request) (*URI, error)

func ProxyURI(fixedURI *URI) Proxy { _ = "STUB: not implemented"; return *new(Proxy) }
