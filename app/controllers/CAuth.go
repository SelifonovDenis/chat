package controllers

import (
	"github.com/revel/revel"
	"just_chatting/app/modules/auth"
	"just_chatting/app/modules/entity"
)

type CAuth struct {
	*revel.Controller
	provider *auth.AuthProvider
	response entity.Response
}

func (c *CAuth) Init() (err error) {
	c.provider, err = auth.NewAuthProvider()
	if err != nil{
		return
	}
	c.response = entity.Response{}
	return
}

//CAuth.Login Функция входа в чат по никнейму
func (c *CAuth) Login() revel.Result {
	err := c.Init()
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	user := &entity.User{}
	err = c.Params.BindJSON(&user)
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	user, err = c.provider.Login(user)
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	err = c.Session.Set("user", user)
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	c.response.Type = entity.ResponseTypeData
	c.response.Data = user
	return c.RenderJSON(c.response)
}

//CAuth.Logout Функция выхода из чата
func (c *CAuth) Logout() revel.Result {
	err := c.Init()
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	user := &entity.User{}
	err = c.Params.BindJSON(&user)
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	err = c.provider.Logout(user)
	if err != nil {
		c.response.Type = entity.ResponseTypeError
		c.response.Data = err.Error()
		return c.RenderJSON(c.response)
	}

	c.Session.Del("user")
	c.response.Type = entity.ResponseTypeData
	c.response.Data = true
	return c.RenderJSON(c.response)
}
