package chat

import (
	"html"
	"just_chatting/app/modules/entity"
	"time"
)

type ChatProvider struct {
	mapper *ChatMapper
}

func NewChatProvider() (*ChatProvider, error) {
	mapper, err := NewChatMapper()
	if err!=nil {
		return &ChatProvider{}, err
	}
	return &ChatProvider{mapper: mapper}, nil
}

func (p *ChatProvider) GetMessages() ([]*entity.Message, error){
	messages, err := p.mapper.GetMessages()
	if err != nil{
		return nil, err
	}

	return messages, nil
}

func (p *ChatProvider) NewMessage(message *entity.Message) error{
	p.PrepareMessage(message)
	err := p.mapper.NewMessage(message)
	if err != nil{
		return err
	}

	return nil
}

func (p *ChatProvider) PrepareMessage(message *entity.Message){
	message.SendlerId = message.Sendler.Id
	message.Time = time.Now()
	message.Text = html.EscapeString(message.Text)
}
