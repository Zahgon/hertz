package utils

import (
	errs "github.com/cloudwego/hertz/pkg/common/errors"
)

var errNeedMore = errs.New(errs.ErrNeedMore, errs.ErrorTypePublic, "cannot find trailing lf")

func Assert(guard bool, text string) { _ = "STUB: not implemented"; return }

type H map[string]interface{}

func IsTrueString(str string) bool { _ = "STUB: not implemented"; return false }

func NameOfFunction(f interface{}) string { _ = "STUB: not implemented"; return "" }

func CaseInsensitiveCompare(a, b []byte) bool { _ = "STUB: not implemented"; return false }

func NormalizeHeaderKey(b []byte, disableNormalizing bool) { _ = "STUB: not implemented"; return }

func NextLine(b []byte) ([]byte, []byte, error) { _ = "STUB: not implemented"; return nil, nil, nil }

func FilterContentType(content string) string { _ = "STUB: not implemented"; return "" }
