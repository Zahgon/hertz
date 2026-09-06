package retry

import (
	"time"
)

type Config struct {
	MaxAttemptTimes uint

	Delay time.Duration

	MaxDelay time.Duration

	MaxJitter time.Duration

	DelayPolicy DelayPolicyFunc
}

func (o *Config) Apply(opts []Option) { _ = "STUB: not implemented"; return }

type DelayPolicyFunc func(attempts uint, err error, retryConfig *Config) time.Duration

func DefaultDelayPolicy(_ uint, _ error, _ *Config) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func FixedDelayPolicy(_ uint, _ error, retryConfig *Config) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func RandomDelayPolicy(_ uint, _ error, retryConfig *Config) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func BackOffDelayPolicy(attempts uint, _ error, retryConfig *Config) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func CombineDelay(delays ...DelayPolicyFunc) DelayPolicyFunc {
	_ = "STUB: not implemented"
	return *new(DelayPolicyFunc)
}

func Delay(attempts uint, err error, retryConfig *Config) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
