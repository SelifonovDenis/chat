package controllers

import (
	"github.com/revel/revel"
	"just_chatting/app/chatroom"
	"just_chatting/app/modules/chat"
	"just_chatting/app/modules/entity"
)

type CChat struct {
	*revel.Controller
	provider *chat.ChatProvider
	response entity.Response
}

func (c CChat) ChatSocket(ws revel.ServerWebSocket) revel.Result {
	// проверка вебсокета.
	if ws == nil {
		return nil
	}

	var err error
	c.provider, err = chat.NewChatProvider()
	if err != nil{
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	messages, err :=c.provider.GetMessages()
	if err!=nil{
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	for _, message := range messages {
		if ws.MessageSendJSON(&message) != nil {
			return nil
		}
	}
	subscription := chatroom.Subscribe()
	//получение сообщений с фронта
	go func() {
		var msg *entity.Message
		for {
			err := ws.MessageReceiveJSON(&msg)
			if err != nil {
				return
			}
			err = c.provider.NewMessage(msg)
			if err!=nil{
				return
			}
			chatroom.Say(*msg)
		}
	}()

	// проверка на новые сообщения
	for {
		select {
		case event := <-subscription.New:
			if ws.MessageSendJSON(&event) != nil {
				// They disconnected.
				return nil
			}
		}
	}

}

