package chat

import (
	"gorm.io/gorm"
	"just_chatting/app/modules/entity"
	"just_chatting/app/modules/helpers"
)

type ChatMapper struct {
	db *gorm.DB
}

func NewChatMapper() (*ChatMapper, error) {
	db, err := helpers.ConnectToDB()
	if err != nil {
		return &ChatMapper{}, err
	}
	return &ChatMapper{db: db}, nil
}

//
func (m *ChatMapper) GetMessages() ([]*entity.Message, error) {
	var messages []*entity.Message
	result := m.db.Joins("Sendler").
		Order("time asc").
		Limit(50).
		Offset(0).
		Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

//
func (m *ChatMapper) NewMessage(message *entity.Message) error {
	result := m.db.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

