package protocol

type Trailer struct {
	h                  []argsKV
	bufKV              argsKV
	disableNormalizing bool
}

func (t *Trailer) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (t *Trailer) Peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (t *Trailer) Del(key string) { _ = "STUB: not implemented"; return }

func (t *Trailer) VisitAll(f func(key, value []byte)) { _ = "STUB: not implemented"; return }

func (t *Trailer) Set(key, value string) error { _ = "STUB: not implemented"; return nil }

func (t *Trailer) Add(key, value string) error { _ = "STUB: not implemented"; return nil }

func (t *Trailer) addArgBytes(key, value []byte, noValue bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Trailer) setArgBytes(key, value []byte, noValue bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Trailer) UpdateArgBytes(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (t *Trailer) GetTrailers() []argsKV { _ = "STUB: not implemented"; return nil }

func (t *Trailer) Empty() bool { _ = "STUB: not implemented"; return false }

func (t *Trailer) GetBytes() []byte { _ = "STUB: not implemented"; return nil }

func (t *Trailer) ResetSkipNormalize() { _ = "STUB: not implemented"; return }

func (t *Trailer) Reset() { _ = "STUB: not implemented"; return }

func (t *Trailer) DisableNormalizing() { _ = "STUB: not implemented"; return }

func (t *Trailer) IsDisableNormalizing() bool { _ = "STUB: not implemented"; return false }

func (t *Trailer) CopyTo(dst *Trailer) { _ = "STUB: not implemented"; return }

func (t *Trailer) SetTrailers(trailers []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (t *Trailer) Header() []byte { _ = "STUB: not implemented"; return nil }

func (t *Trailer) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func IsBadTrailer(key []byte) bool { _ = "STUB: not implemented"; return false }
