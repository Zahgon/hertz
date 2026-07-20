package consts

import "time"

const (
	DefaultMaxInMemoryFileSize = 16 * 1024 * 1024

	DefaultDialTimeout = time.Second

	DefaultMaxConnsPerHost = 512

	DefaultMaxIdleConnDuration = 10 * time.Second

	DefaultMaxIdempotentCallAttempts = 1

	DefaultMaxRetryTimes = 1
)
