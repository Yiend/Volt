package server

import (
	"fmt"
	"github.com/kataras/iris"
	"os"
	"volt/app/identity"
	"volt/app/router"
	"volt/app/startup"
)

func Run()  {
	fmt.Printf("服务启动，进程号：%d/n",os.Getpid())

    app :=startupApp()
	app.Run(iris.Addr(":8021"),iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  true,
		FireMethodNotAllowed:              true,
		DisableBodyConsumptionOnUnmarshal: true,
		DisableAutoFireStatusCode:         true,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))

}


func startupApp() *startup.Startup {
	app := startup.New("VoltApp", "wirehomedev@gmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, router.Configure)
	return app
}
