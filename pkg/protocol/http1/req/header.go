package req

import (
	"errors"

	"github.com/cloudwego/hertz/internal/bytesconv"
	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var errEOFReadHeader = errs.NewPublic("error when reading request headers: EOF")

func WriteHeader(h *protocol.RequestHeader, w network.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadHeader(h *protocol.RequestHeader, r network.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadHeaderWithLimit(h *protocol.RequestHeader, r network.Reader, maxHeaderBytes int) error {
	_ = "STUB: not implemented"
	return nil
}

func tryReadWithLimit(h *protocol.RequestHeader, r network.Reader, n, maxHeaderBytes int) error {
	_ = "STUB: not implemented"
	return nil
}

func parse(h *protocol.RequestHeader, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

const (
	maxCheckMethodLen = 10

	validMethodCharTable = bytesconv.ValidHeaderFieldNameTable
)

var errMalformedHTTPRequest = errors.New("malformed HTTP request")

var errBothTEAndCL = errors.New("both Transfer-Encoding and Content-Length headers are present in request: potential HTTP request smuggling")

var errDuplicateCL = errors.New("duplicate Content-Length header with conflicting values")

var errUnsupportedTE = errors.New("unsupported Transfer-Encoding: only \"chunked\" is accepted")

var errMultipleTE = errors.New("multiple Transfer-Encoding headers are not allowed")

func parseFirstLine(h *protocol.RequestHeader, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func validHeaderFieldValue(val []byte) bool { _ = "STUB: not implemented"; return false }

func parseHeaders(h *protocol.RequestHeader, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
