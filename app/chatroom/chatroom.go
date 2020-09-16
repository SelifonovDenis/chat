package chatroom

import (
	"container/list"
	"just_chatting/app/modules/entity"
)

func init(){
	go chatroom()
}

var (
	// Каналы
	// Подписки
	subscribe = make(chan (chan <- Subscription))
	// Отписки
	unsubscribe = make(chan (<-chan entity.Message))
	// отправка сообщений
	publish = make(chan entity.Message)
)

//подписка на события
type Subscription struct {
	New     <- chan entity.Message // свой канал для прослушки новых сообщений.
}

// новая подписка
func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}

//отписка
func (s Subscription) Cancel() {
	unsubscribe <- s.New
	drain(s.New)
}

//передача сообщения
func Say(message entity.Message) {
	publish <- message
}

// горутина отправки сообщений подписчикам
func chatroom() {
	subscribers := list.New()

	for {
		select {
		case ch := <-subscribe:
			subscriber := make(chan entity.Message, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{subscriber}

		case event := <-publish:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan entity.Message) <- event
			}

		case unsub := <-unsubscribe:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				if ch.Value.(chan entity.Message) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		}
	}
}

// Helpers
func drain(ch <- chan entity.Message) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}