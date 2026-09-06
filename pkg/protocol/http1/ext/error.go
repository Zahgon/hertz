package ext

import (
	errs "github.com/cloudwego/hertz/pkg/common/errors"
)

var (
	errNeedMore     = errs.New(errs.ErrNeedMore, errs.ErrorTypePublic, "cannot find trailing lf")
	errBodyTooLarge = errs.New(errs.ErrBodyTooLarge, errs.ErrorTypePublic, "ext")
)

func HeaderError(typ string, err, errParse error, b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func headerErrorMsg(typ string, err error, b []byte) error { _ = "STUB: not implemented"; return nil }
