package eventbus

import (
	"math/rand"
	"strconv"
)

const (
	defaultBufferLength = 8048
)

type Handler func(any) error

type Topic string

type SubscriberID string

func newSubscriberID(consumerName string) SubscriberID {
	return SubscriberID(consumerName + "-" + strconv.Itoa(rand.Int())) //nolint:gosec
}
