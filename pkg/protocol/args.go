package protocol

import (
	"io"

	"github.com/cloudwego/hertz/internal/nocopy"
)

const (
	argsNoValue  = true
	ArgsHasValue = false
)

type argsScanner struct {
	b []byte
}

type Args struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	args []argsKV
	buf  []byte
}

func (a *Args) Set(key, value string) { _ = "STUB: not implemented"; return }

func (a *Args) Reset() { _ = "STUB: not implemented"; return }

func (a *Args) CopyTo(dst *Args) { _ = "STUB: not implemented"; return }

func (a *Args) Del(key string) { _ = "STUB: not implemented"; return }

func (a *Args) DelBytes(key []byte) { _ = "STUB: not implemented"; return }

func (s *argsScanner) next(kv *argsKV) bool { _ = "STUB: not implemented"; return false }

func decodeArgAppend(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func allocArg(h []argsKV) ([]argsKV, *argsKV) { _ = "STUB: not implemented"; return nil, nil }

func releaseArg(h []argsKV) []argsKV { _ = "STUB: not implemented"; return nil }

func updateArgBytes(h []argsKV, key, value []byte) []argsKV { _ = "STUB: not implemented"; return nil }

func setArgBytes(h []argsKV, key, value []byte, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func setArg(h []argsKV, key, value string, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func peekArgBytes(h []argsKV, k []byte) []byte { _ = "STUB: not implemented"; return nil }

func peekAllArgBytesToDst(dst [][]byte, h []argsKV, k []byte) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func delAllArgsBytes(args []argsKV, key []byte) []argsKV { _ = "STUB: not implemented"; return nil }

func delAllArgs(args []argsKV, key string) []argsKV { _ = "STUB: not implemented"; return nil }

func (a *Args) Has(key string) bool { _ = "STUB: not implemented"; return false }

func hasArg(h []argsKV, key string) bool { _ = "STUB: not implemented"; return false }

func (a *Args) String() string { _ = "STUB: not implemented"; return "" }

func decodeArgAppendNoPlus(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func peekArgStr(h []argsKV, k string) []byte { _ = "STUB: not implemented"; return nil }

func peekArgStrExists(h []argsKV, k string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (a *Args) QueryString() []byte { _ = "STUB: not implemented"; return nil }

func (a *Args) ParseBytes(b []byte) { _ = "STUB: not implemented"; return }

func (a *Args) Peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (a *Args) PeekExists(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (a *Args) PeekAll(key string) [][]byte { _ = "STUB: not implemented"; return nil }

func visitArgs(args []argsKV, f func(k, v []byte)) { _ = "STUB: not implemented"; return }

func (a *Args) Len() int { _ = "STUB: not implemented"; return 0 }

func (a *Args) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (a *Args) VisitAll(f func(key, value []byte)) { _ = "STUB: not implemented"; return }

func (a *Args) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (a *Args) Add(key, value string) { _ = "STUB: not implemented"; return }
