package utils

import (
	"io"

	"github.com/cloudwego/hertz/pkg/network"
)

func CopyBuffer(dst network.Writer, src io.Reader, buf []byte) (written int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func copyBuffer(dst network.Writer, src io.Reader, buf []byte) (written int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func CopyZeroAlloc(w network.Writer, r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
