package ext

import (
	errs "github.com/cloudwego/hertz/pkg/common/errors"
)

var errInvalidName = errs.NewPublic("invalid header name")

type HeaderScanner struct {
	B     []byte
	Key   []byte
	Value []byte
	Err   error

	HLen int

	DisableNormalizing bool

	nextColon   int
	nextNewLine int

	initialized bool
}

type HeaderValueScanner struct {
	B     []byte
	Value []byte
}

func (s *HeaderScanner) Next() bool { _ = "STUB: not implemented"; return false }

func (s *HeaderValueScanner) next() bool { _ = "STUB: not implemented"; return false }

func HasHeaderValue(s, value []byte) bool { _ = "STUB: not implemented"; return false }
