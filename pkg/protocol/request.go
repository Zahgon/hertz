package protocol

import (
	"io"
	"mime/multipart"
	"net/url"
	"sync"

	"github.com/cloudwego/hertz/internal/nocopy"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/errors"
)

var (
	ErrMissingFile = errors.NewPublic("http: no such file")

	responseBodyPool bytebufferpool.Pool
	requestBodyPool  bytebufferpool.Pool

	requestPool sync.Pool
)

var NoBody = noBody{}

type noBody struct{}

func (noBody) Read([]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (noBody) Close() error             { _ = "STUB: not implemented"; return nil }

type Request struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	Header RequestHeader

	uri      URI
	postArgs Args

	bodyStream      io.Reader
	w               requestBodyWriter
	body            *bytebufferpool.ByteBuffer
	bodyRaw         []byte
	maxKeepBodySize int

	multipartForm         *multipart.Form
	multipartFormBoundary string

	parsedURI      bool
	parsedPostArgs bool

	isTLS bool

	multipartFiles  []*File
	multipartFields []*MultipartField

	options *config.RequestOptions
}

type requestBodyWriter struct {
	r *Request
}

type File struct {
	Name      string
	ParamName string
	io.Reader
}

type MultipartField struct {
	Param       string
	FileName    string
	ContentType string
	io.Reader
}

func (w *requestBodyWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (req *Request) Options() *config.RequestOptions { _ = "STUB: not implemented"; return nil }

func (req *Request) AppendBody(p []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck

func (req *Request) BodyBuffer() *bytebufferpool.ByteBuffer { _ = "STUB: not implemented"; return nil }

func (req *Request) MayContinue() bool { _ = "STUB: not implemented"; return false }

func (req *Request) Scheme() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) resetSkipHeaderAndConn() { _ = "STUB: not implemented"; return }

func (req *Request) ResetSkipHeader() { _ = "STUB: not implemented"; return }

func SwapRequestBody(a, b *Request) { _ = "STUB: not implemented"; return }

func (req *Request) Reset() { _ = "STUB: not implemented"; return }

func (req *Request) IsURIParsed() bool { _ = "STUB: not implemented"; return false }

func (req *Request) PostArgString() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) MultipartForm() (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *Request) AppendBodyString(s string) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck

func (req *Request) SetRequestURI(requestURI string) { _ = "STUB: not implemented"; return }

func (req *Request) SetMaxKeepBodySize(n int) { _ = "STUB: not implemented"; return }

func (req *Request) RequestURI() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) FormFile(name string) (*multipart.FileHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *Request) SetHost(host string) { _ = "STUB: not implemented"; return }

func (req *Request) Host() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) SetIsTLS(isTLS bool) { _ = "STUB: not implemented"; return }

func (req *Request) SwapBody(body []byte) []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (req *Request) CopyTo(dst *Request) { _ = "STUB: not implemented"; return }

func (req *Request) CopyToSkipBody(dst *Request) { _ = "STUB: not implemented"; return }

func (req *Request) BodyBytes() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) ResetBody() { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (req *Request) SetBodyRaw(body []byte) { _ = "STUB: not implemented"; return }

func (req *Request) SetMultipartFormBoundary(b string) { _ = "STUB: not implemented"; return }

func (req *Request) MultipartFormBoundary() string { _ = "STUB: not implemented"; return "" }

func (req *Request) SetBody(body []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (req *Request) SetBodyString(body string) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (req *Request) SetQueryString(queryString string) { _ = "STUB: not implemented"; return }

func (req *Request) SetFormData(data map[string]string) { _ = "STUB: not implemented"; return }

func (req *Request) SetFormDataFromValues(data url.Values) { _ = "STUB: not implemented"; return }

func (req *Request) SetFile(param, filePath string) { _ = "STUB: not implemented"; return }

func (req *Request) SetFiles(files map[string]string) { _ = "STUB: not implemented"; return }

func (req *Request) SetFileReader(param, fileName string, reader io.Reader) {
	_ = "STUB: not implemented"
	return
}

func (req *Request) SetMultipartFormData(data map[string]string) { _ = "STUB: not implemented"; return }

func (req *Request) MultipartFiles() []*File { _ = "STUB: not implemented"; return nil }

func (req *Request) SetMultipartField(param, fileName, contentType string, reader io.Reader) {
	_ = "STUB: not implemented"
	return
}

func (req *Request) SetMultipartFields(fields ...*MultipartField) {
	_ = "STUB: not implemented"
	return
}

func (req *Request) MultipartFields() []*MultipartField { _ = "STUB: not implemented"; return nil }

func (req *Request) SetBasicAuth(username, password string) { _ = "STUB: not implemented"; return }

func (req *Request) BasicAuth() (username, password string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

var prefix = []byte{'B', 'a', 's', 'i', 'c', ' '}

func parseBasicAuth(auth []byte) (username, password string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (req *Request) SetAuthToken(token string) { _ = "STUB: not implemented"; return }

func (req *Request) SetAuthSchemeToken(scheme, token string) { _ = "STUB: not implemented"; return }

func (req *Request) SetHeader(header, value string) { _ = "STUB: not implemented"; return }

func (req *Request) SetHeaders(headers map[string]string) { _ = "STUB: not implemented"; return }

func (req *Request) SetCookie(key, value string) { _ = "STUB: not implemented"; return }

func (req *Request) SetCookies(hc map[string]string) { _ = "STUB: not implemented"; return }

func (req *Request) SetMethod(method string) { _ = "STUB: not implemented"; return }

func (req *Request) OnlyMultipartForm() bool { _ = "STUB: not implemented"; return false }

func (req *Request) HasMultipartForm() bool { _ = "STUB: not implemented"; return false }

func (req *Request) IsBodyStream() bool { _ = "STUB: not implemented"; return false }

func (req *Request) BodyStream() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (req *Request) SetBodyStream(bodyStream io.Reader, bodySize int) {
	_ = "STUB: not implemented"
	return
}

func (req *Request) ConstructBodyStream(body *bytebufferpool.ByteBuffer, bodyStream io.Reader) {
	_ = "STUB: not implemented"
	return
}

func (req *Request) BodyWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (req *Request) PostArgs() *Args { _ = "STUB: not implemented"; return nil }

func (req *Request) parsePostArgs() { _ = "STUB: not implemented"; return }

func (req *Request) BodyE() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:errcheck

func (req *Request) Body() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) BodyWriteTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (req *Request) CloseBodyStream() error { _ = "STUB: not implemented"; return nil }

func (req *Request) URI() *URI { _ = "STUB: not implemented"; return nil }

func (req *Request) ParseURI() { _ = "STUB: not implemented"; return }

func (req *Request) RemoveMultipartFormFiles() { _ = "STUB: not implemented"; return }

//nolint:errcheck

func AddMultipartFormField(w *multipart.Writer, mf *MultipartField) error {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) Method() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) Path() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) QueryString() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) SetOptions(opts ...config.RequestOption) { _ = "STUB: not implemented"; return }

func (req *Request) ConnectionClose() bool { _ = "STUB: not implemented"; return false }

func (req *Request) SetConnectionClose() { _ = "STUB: not implemented"; return }

func (req *Request) ResetWithoutConn() { _ = "STUB: not implemented"; return }

func AcquireRequest() *Request { _ = "STUB: not implemented"; return nil }

func ReleaseRequest(req *Request) { _ = "STUB: not implemented"; return }

func NewRequest(method, url string, body io.Reader) *Request { _ = "STUB: not implemented"; return nil }
