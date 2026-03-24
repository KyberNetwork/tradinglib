package eventbus

import (
	"context"
	"sync"

	"github.com/cskr/pubsub"
)

type Manager struct {
	rw   sync.RWMutex
	subs map[Topic]Subscription
}

func NewManager() *Manager {
	return &Manager{
		rw:   sync.RWMutex{},
		subs: make(map[Topic]Subscription),
	}
}

type Subscription struct {
	topic      string
	ps         *pubsub.PubSub
	subscriber map[SubscriberID]chan any
}

// StartConsume subscribe with a handler.
func (m *Manager) StartConsume(ctx context.Context, consumerName string, topic Topic, fn Handler) {
	id, c := m.Subscribe(consumerName, topic)

	for {
		select {
		case <-ctx.Done():
			m.Unsubscribe(topic, id)

			return

		case v, ok := <-c:
			if !ok {
				continue
			}

			// Ignore error here, it should be handled in handler.
			_ = fn(v)
		}
	}
}

// StartConsumeMultiple subscribe with a handler.
func (m *Manager) StartConsumeMultiple(ctx context.Context, consumerName string, topic Topic, fn Handler, worker int) {
	id, c := m.Subscribe(consumerName, topic)

	wg := sync.WaitGroup{}
	defer wg.Wait()

	for i := 0; i < worker; i++ { //nolint:modernize
		wg.Add(1)

		go func() { //nolint:modernize
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					m.Unsubscribe(topic, id)

					return

				case v, ok := <-c:
					if !ok {
						continue
					}

					// Ignore error here, it should be handled in handler.
					_ = fn(v)
				}
			}
		}()
	}
}

func (m *Manager) Publish(topic Topic, message any) {
	sub, exist := m.getSub(topic)
	if !exist {
		m.rw.Lock()
		{
			sub = Subscription{
				topic:      string(topic),
				ps:         pubsub.New(defaultBufferLength),
				subscriber: make(map[SubscriberID]chan any),
			}

			m.subs[topic] = sub
		}

		m.rw.Unlock()
	}

	sub.ps.Pub(message, sub.topic)
}

func (m *Manager) Subscribe(consumerName string, topic Topic) (SubscriberID, <-chan any) {
	m.rw.Lock()
	defer m.rw.Unlock()

	subscription, exist := m.subs[topic]
	if !exist {
		subscription = Subscription{
			topic:      string(topic),
			ps:         pubsub.New(defaultBufferLength),
			subscriber: make(map[SubscriberID]chan any),
		}

		m.subs[topic] = subscription
	}

	id := newSubscriberID(consumerName)
	c := subscription.ps.Sub(subscription.topic)
	subscription.subscriber[id] = c

	return id, c
}

func (m *Manager) Unsubscribe(topic Topic, id SubscriberID) {
	m.rw.Lock()
	defer m.rw.Unlock()

	sub, exist := m.subs[topic]
	if !exist {
		return
	}

	c, exist := sub.subscriber[id]
	if !exist {
		return
	}

	sub.ps.Unsub(c, sub.topic)
	delete(sub.subscriber, id)
}

func (m *Manager) getSub(topic Topic) (Subscription, bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	s, exist := m.subs[topic]

	return s, exist
}

func (m *Manager) GetStats() map[string]map[string]int64 {
	m.rw.RLock()
	defer m.rw.RUnlock()

	stats := make(map[string]map[string]int64)

	for topic, sub := range m.subs {
		stats[string(topic)] = make(map[string]int64)
		for subID, c := range sub.subscriber {
			stats[string(topic)][string(subID)] = int64(len(c))
		}
	}

	return stats
}
