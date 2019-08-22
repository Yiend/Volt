package identity

import (
	"github.com/kataras/iris"
	"time"
	"volt/app/startup"
)

func New(s *startup.Startup) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		ctx.Header("App-Name", s.AppName)
		ctx.Header("App-Owner", s.AppOwner)
		ctx.Header("App-Since", time.Since(s.AppSpawnDate).String())

		ctx.Header("Server", "volt")

		ctx.ViewData("AppName", s.AppName)
		ctx.ViewData("AppOwner", s.AppOwner)
		ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(s *startup.Startup) {
	h := New(s)
	s.UseGlobal(h)
}
