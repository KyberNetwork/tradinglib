package mev

import (
	"golang.org/x/time/rate"
)

type NewBundleSendleClientOption func(newBundleSendleClientOptions) newBundleSendleClientOptions

type newBundleSendleClientOptions struct {
	enableSendPrivateRaw    bool
	builderNetRefundAddress string
	sendBundleLimiter       *rate.Limiter
}

func WithSendPrivateRaw() NewBundleSendleClientOption {
	return func(opt newBundleSendleClientOptions) newBundleSendleClientOptions {
		opt.enableSendPrivateRaw = true

		return opt
	}
}

func WithBuilderNetRefundAddress(addr string) NewBundleSendleClientOption {
	return func(opt newBundleSendleClientOptions) newBundleSendleClientOptions {
		opt.builderNetRefundAddress = addr
		return opt
	}
}

// WithSendBundleRateLimit caps the eth_sendBundle requests that one Client sends.
// A client over the limit fails the call at once instead of waiting, because a late
// bundle has no value. A non-positive requestsPerSecond or burst leaves the client
// unlimited, which matches the zero value of the option.
func WithSendBundleRateLimit(requestsPerSecond float64, burst int) NewBundleSendleClientOption {
	return func(opt newBundleSendleClientOptions) newBundleSendleClientOptions {
		if requestsPerSecond <= 0 || burst <= 0 {
			return opt
		}
		opt.sendBundleLimiter = rate.NewLimiter(rate.Limit(requestsPerSecond), burst)

		return opt
	}
}
