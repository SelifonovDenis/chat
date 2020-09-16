package controllers

import (
	"github.com/revel/revel"
	"just_chatting/app/modules/entity"
)

type CUser struct {
	*revel.Controller
	response entity.Response
}
// получение пользователя из сессии
func (c *CUser) GetUser() revel.Result {

	user, err:= c.Session.Get("user")
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	} else {
		if c.Session.SessionTimeoutExpiredOrMissing(){
			c.response.Type = entity.ResponseTypeError
			c.response.Data = "Сессия истекла"
			return c.RenderJSON(c.response)
		}
	}
	c.response.Type = entity.ResponseTypeData
	c.response.Data = user
	return c.RenderJSON(c.response)
}