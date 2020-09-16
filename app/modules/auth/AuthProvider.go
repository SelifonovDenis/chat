package auth

import (
	"fmt"
	"github.com/revel/revel"
	"html"
	"just_chatting/app/modules/entity"
	"time"
)

type AuthProvider struct {
	mapper *AuthMapper
}

func NewAuthProvider() (*AuthProvider, error) {
	mapper, err := NewAuthMapper()
	if err!=nil {
		return &AuthProvider{}, err
	}
	return &AuthProvider{mapper: mapper}, nil
}

func (p *AuthProvider) Login(user *entity.User) (*entity.User, error) {
	err := p.PrepareUser(user)
	if err!=nil{
		return user, err
	}
	// проверка никнейма на существование
	user, err = p.mapper.CheckNickname(user)
	if err != nil {
		return user, err
	}
	// Создание нового пользователя
	user,err = p.mapper.NewUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *AuthProvider) Logout(user *entity.User) error {
	// обнуляем время
	user.DateCreate = time.Time{}
	err := p.mapper.Logout(user)
	if err!=nil{
		return err
	}
	return nil
}

func (p *AuthProvider) PrepareUser(user *entity.User) (err error){
	// получаем таймаут сессии
	expires, found:= revel.Config.String("session.expires")
	if !found{
		return fmt.Errorf("Не найден таймаут сессси.")
	}
	duration, err := time.ParseDuration(expires)
	if err!=nil{
		return err
	}
	// изменяем дату для проверки активности пользователя
	user.DateCreate = time.Now().Add(-duration)
	// xxs защита
	user.Nickname = html.EscapeString(user.Nickname)
	return nil
}
