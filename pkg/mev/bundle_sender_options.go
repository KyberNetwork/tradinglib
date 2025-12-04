package mev

type NewBundleSendleClientOption func(newBundleSendleClientOptions) newBundleSendleClientOptions

type newBundleSendleClientOptions struct {
	enableSendPrivateRaw    bool
	builderNetRefundAddress string
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
