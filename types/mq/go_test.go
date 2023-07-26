package mq

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBrokerSend(t *testing.T) {

	broker := GetBroker()

	go func() {
		for i := 0; ; i++ {
			err := broker.Send(Msg{Content: fmt.Sprintf("%v", i)})
			if err != nil {
				return
			}
		}
	}()

	wg := sync.WaitGroup{}
	for index := 0; index < 3; index++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			consumeGroup, err := broker.Subscribe(100)
			if err != nil {
				return
			}
			for {
				select {
				case m, ok := <-consumeGroup:
					t.Log("group", index, m)
					if !ok {
						return
					}
				}

			}

		}(index)
		time.Sleep(time.Second)
	}
	time.Sleep(10 * time.Second)
	broker.Close()
	wg.Wait()
}
