package filter

import (
	"github.com/revel/revel"
	"strings"
)

// фильтр проверки авторизации
func AuthFilter(c *revel.Controller, fc []revel.Filter) {
	allowedPathList := []string{"/", "/favicon.ico", "/login"}
	allowed := false
	for _, path := range allowedPathList {
		if path == c.Request.URL.Path {
			allowed = true
		}
		if strings.HasPrefix(c.Request.URL.Path, "/public/dist"){
			allowed = true
		}
	}
	if !allowed {
		_, err:= c.Session.Get("user")
		if err != nil {
			c.Result = c.Redirect("/")
		} else {
			if c.Session.SessionTimeoutExpiredOrMissing(){
				c.Session.Del("user")
				c.Result = c.Redirect("/")
			}
		}
	}
	fc[0](c, fc[1:])
}
