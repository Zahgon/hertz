//go:build windows

package protocol

func addLeadingSlash(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func checkSchemeWhenCharIsColon(i int, rawURL []byte) (scheme, path []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}
