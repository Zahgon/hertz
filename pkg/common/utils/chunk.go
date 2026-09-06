package utils

import (
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
)

var errBrokenChunk = errors.NewPublic("cannot find crlf at the end of chunk")

func ParseChunkSize(r network.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func SkipCRLF(reader network.Reader) error { _ = "STUB: not implemented"; return nil }
