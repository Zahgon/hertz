package protocol

import (
	"io"
	"mime/multipart"
	"net/textproto"
)

func ReadMultipartForm(r io.Reader, boundary string, size, maxInMemoryFileSize int) (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteMultipartForm(w io.Writer, f *multipart.Form, boundary string) error {
	_ = "STUB: not implemented"
	return nil
}

func MarshalMultipartForm(f *multipart.Form, boundary string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteMultipartFormFile(w *multipart.Writer, fieldName, fileName string, r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateMultipartHeader(param, fileName, contentType string) textproto.MIMEHeader {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader)
}

func AddFile(w *multipart.Writer, fieldName, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func ParseMultipartForm(r io.Reader, request *Request, size, maxInMemoryFileSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func SetMultipartFormWithBoundary(req *Request, m *multipart.Form, boundary string) {
	_ = "STUB: not implemented"
	return
}
