package eventbus_test

import (
	"context"
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/eventbus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test1(t *testing.T) {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)

	var (
		m                    = eventbus.NewManager()
		topic eventbus.Topic = "hello"
		msg                  = "world"
	)

	id, c := m.Subscribe("", topic)
	t.Log("id", id)
	m.Publish(topic, msg)

	s := <-c
	t.Log("s", s)
}

func Test2(t *testing.T) {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)

	var (
		m                          = eventbus.NewManager()
		topic       eventbus.Topic = "hello"
		msg                        = "world"
		ctx, cancel                = context.WithCancel(context.Background())
	)

	go m.StartConsume(
		ctx,
		"",
		topic,
		func(i any) error {
			s, ok := i.(string)
			assert.True(t, ok, i)

			assert.Equal(t, msg, s)
			cancel()

			return nil
		})

	time.Sleep(time.Second)
	m.Publish(topic, msg)
	<-ctx.Done()
}
