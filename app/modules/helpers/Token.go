package helpers

import (
	"just_chatting/app/modules/entity"
	"strconv"
	"strings"
)

type TokenHelper struct {
	user *entity.User
}

func NewTokenHelper() TokenHelper {
	return TokenHelper{
		user: nil,
	}
}

// метод создания простейшего токена
func (t *TokenHelper) CreateToken() (token string, err error) {
	token = strconv.Itoa(t.user.Id) + ":" + t.user.Nickname
	return
}

func (t *TokenHelper) CheckToken(token string) (result bool, err error) {
	userData := strings.Split(token, ":")
	if len(userData) != 2 {
		result = false
	}
	//проверка авторизации
	return
}
