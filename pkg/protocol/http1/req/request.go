package req

import (
	"fmt"

	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var (
	errRequestHostRequired = errs.NewPublic("missing required Host header in request")
	errGetOnly             = errs.NewPublic("non-GET request received")
	errBodyTooLarge        = errs.New(errs.ErrBodyTooLarge, errs.ErrorTypePublic, "http1/req")
	errHeaderTooLarge      = errs.New(errs.ErrHeaderTooLarge, errs.ErrorTypePublic, "http1/req")
)

type h1Request struct {
	*protocol.Request
}

func (h1Req *h1Request) String() string { _ = "STUB: not implemented"; return "" }

func GetHTTP1Request(req *protocol.Request) fmt.Stringer {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer)
}

func ReadHeaderAndLimitBody(req *protocol.Request, r network.Reader, maxBodySize int, preParse ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func Read(req *protocol.Request, r network.Reader, preParse ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func Write(req *protocol.Request, w network.Writer) error { _ = "STUB: not implemented"; return nil }

func ProxyWrite(req *protocol.Request, w network.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func write(req *protocol.Request, w network.Writer, usingProxy bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func ContinueReadBodyStream(req *protocol.Request, zr network.Reader, maxBodySize int, preParseMultipartForm ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ContinueReadBody(req *protocol.Request, r network.Reader, maxBodySize int, preParseMultipartForm ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadBodyStream(req *protocol.Request, zr network.Reader, maxBodySize int, getOnly, preParseMultipartForm bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadLimitBody(req *protocol.Request, r network.Reader, maxBodySize int, getOnly, preParseMultipartForm bool) error {
	_ = "STUB: not implemented"
	return nil
}

func writeBodyStream(req *protocol.Request, w network.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func handleMultipart(req *protocol.Request) error { _ = "STUB: not implemented"; return nil }
