package adaptor

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/protocol"
)

func GetCompatRequest(req *protocol.Request) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CopyToHertzRequest(req *http.Request, hreq *protocol.Request) error {
	_ = "STUB: not implemented"
	return nil
}
