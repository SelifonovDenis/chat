package auth

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"just_chatting/app/modules/entity"
	"just_chatting/app/modules/helpers"
)

type AuthMapper struct {
	db *gorm.DB
}

func NewAuthMapper() (*AuthMapper, error) {
	db, err := helpers.ConnectToDB()
	if err != nil {
		return &AuthMapper{}, err
	}
	return &AuthMapper{db: db}, nil
}

// Проверка никнейма
func (m *AuthMapper) CheckNickname(user *entity.User) (*entity.User, error) {
	//пробуем найти активного пользователя с таким ником, у которого не вышел таймаут сессии
	result := m.db.Where("nickname = ? AND date_create > ?", user.Nickname, user.DateCreate).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	// результата нет, продолжаем
	if result.RowsAffected == 0 {
		fmt.Println("rowaffected = 0")
		return user, nil
	}

	return user, errors.New("Такой никнейм уже существует")
}

// Добавление в БД новогопользователя
func (m *AuthMapper) NewUser(user *entity.User) (*entity.User, error) {
	result := m.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// изменяем дату создания пользователя, чтоы сделать его неактивным
func (m *AuthMapper) Logout(user *entity.User) error {
	result := m.db.Save(&user)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}
