package mq

import (
	"errors"
	"sync"
	"time"
)

type Broker struct {
	mutex sync.Mutex
	chans []chan Msg
}

type Msg struct {
	Content string
}

func (b *Broker) Send(m Msg) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, ch := range b.chans {
		select {
		case ch <- m:
		default:
			return errors.New("消息队列已满")

		}
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

func (b *Broker) Subscribe(capacity int) (<-chan Msg, error) {
	res := make(chan Msg, capacity)
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.chans = append(b.chans, res)
	return res, nil
}

func (b *Broker) Close() {
	b.mutex.Lock()
	chans := b.chans
	b.chans = nil
	b.mutex.Unlock()
	for _, ch := range chans {
		close(ch)
	}

}

func GetBroker() *Broker {
	return &Broker{}
}
