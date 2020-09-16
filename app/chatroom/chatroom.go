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

// This function loops forever, handling the chat room pubsub
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


















//// событие для сокета
//type Event struct {
//	Type      string // тип события "join" подключение, "leave" отключение, "message" сообщение
//	User      string
//	Timestamp int    // Unix timestamp (secs)
//	Text      string // What the user said (if Type == "message")
//}
//
////подписка на события
//type Subscription struct {
//	Archive []Event      // Архив событий.
//	New     <-chan Event // Сюда добавляются новые события.
//}
//
//// Owner of a subscription must cancel it when they stop listening to events.
//func (s Subscription) Cancel() {
//	unsubscribe <- s.New // Unsubscribe the channel.
//	drain(s.New)         // Drain it, just in case there was a pending publish.
//}
//
//func newEvent(typ, user, msg string) Event {
//	return Event{typ, user, int(time.Now().Unix()), msg}
//}
//
//// новая подписка
//func Subscribe() Subscription {
//	resp := make(chan Subscription)
//	subscribe <- resp
//	return <-resp
//}
//
//func Join(user string) {
//	publish <- newEvent("join", user, "")
//}
//
//func Say(user, message string) {
//	publish <- newEvent("message", user, message)
//}
//
//func Leave(user string) {
//	publish <- newEvent("leave", user, "")
//}
//
//const archiveSize = 10
//
//var (
//	// Send a channel here to get room events back.  It will send the entire
//	// archive initially, and then new messages as they come in.
//	subscribe = make(chan (chan<- Subscription), 10)
//	// Send a channel here to unsubscribe.
//	unsubscribe = make(chan (<-chan Event), 10)
//	// Send events here to publish them.
//	publish = make(chan Event, 10)
//)
//

