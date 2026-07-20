package mock

import "bytes"

type ExtWriter struct {
	tmp     []byte
	Buf     *bytes.Buffer
	IsFinal *bool
}

func (m *ExtWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (m *ExtWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (m *ExtWriter) Finalize() error { _ = "STUB: not implemented"; return nil }

func (m *ExtWriter) SetBody(body []byte) { _ = "STUB: not implemented"; return }
