package resp

import (
	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var errTimeout = errs.New(errs.ErrTimeout, errs.ErrorTypePublic, "read response header")

func ReadHeader(h *protocol.ResponseHeader, r network.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteHeader(h *protocol.ResponseHeader, w network.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func ConnectionUpgrade(h *protocol.ResponseHeader) bool { _ = "STUB: not implemented"; return false }

func tryRead(h *protocol.ResponseHeader, r network.Reader, n int) error {
	_ = "STUB: not implemented"
	return nil
}

func parseHeaders(h *protocol.ResponseHeader, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parse(h *protocol.ResponseHeader, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parseFirstLine(h *protocol.ResponseHeader, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
