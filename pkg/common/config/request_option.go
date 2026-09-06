package config

import "time"

var preDefinedOpts []RequestOption

type RequestOptions struct {
	tags map[string]string
	isSD bool

	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration

	requestTimeout time.Duration
	start          time.Time
}

type RequestOption struct {
	F func(o *RequestOptions)
}

func NewRequestOptions(opts []RequestOption) *RequestOptions { _ = "STUB: not implemented"; return nil }

func WithTag(k, v string) RequestOption { _ = "STUB: not implemented"; return *new(RequestOption) }

func WithSD(b bool) RequestOption { _ = "STUB: not implemented"; return *new(RequestOption) }

func WithDialTimeout(t time.Duration) RequestOption {
	_ = "STUB: not implemented"
	return *new(RequestOption)
}

func WithReadTimeout(t time.Duration) RequestOption {
	_ = "STUB: not implemented"
	return *new(RequestOption)
}

func WithWriteTimeout(t time.Duration) RequestOption {
	_ = "STUB: not implemented"
	return *new(RequestOption)
}

func WithRequestTimeout(t time.Duration) RequestOption {
	_ = "STUB: not implemented"
	return *new(RequestOption)
}

func (o *RequestOptions) Apply(opts []RequestOption) { _ = "STUB: not implemented"; return }

func (o *RequestOptions) Tag(k string) string { _ = "STUB: not implemented"; return "" }

func (o *RequestOptions) Tags() map[string]string { _ = "STUB: not implemented"; return nil }

func (o *RequestOptions) IsSD() bool { _ = "STUB: not implemented"; return false }

func (o *RequestOptions) DialTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (o *RequestOptions) ReadTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (o *RequestOptions) WriteTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (o *RequestOptions) RequestTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (o *RequestOptions) StartRequest() { _ = "STUB: not implemented"; return }

func (o *RequestOptions) StartTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (o *RequestOptions) CopyTo(dst *RequestOptions) { _ = "STUB: not implemented"; return }

func SetPreDefinedOpts(opts ...RequestOption) { _ = "STUB: not implemented"; return }
