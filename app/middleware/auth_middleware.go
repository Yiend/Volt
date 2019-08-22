package middleware

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"net/url"
	"volt/data/entitys"
)

func Authentication(ctx iris.Context, session *sessions.Session){
	i := session.Get("User")
	user,ok:= i.(*entitys.User)
	if ok && user.ID > 0 {
		ctx.Next()
	}else {
		ctx.Redirect(fmt.Sprintf("/admin/login?returnurl=%s", url.QueryEscape(ctx.Request().RequestURI)))
	}
}

