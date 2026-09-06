package retry

import "time"

type Option struct {
	F func(o *Config)
}

func WithMaxAttemptTimes(maxAttemptTimes uint) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithInitDelay(delay time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxDelay(maxDelay time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDelayPolicy(delayPolicy DelayPolicyFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMaxJitter(maxJitter time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }
