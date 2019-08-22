package startup

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
	"time"
	"volt/app/config"
	"volt/data/entitys"
	"xorm.io/core"
)

type Configurator func(*Startup)

type Startup struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
	Sessions *sessions.Sessions
	Xorm   *xorm.Engine
	AppConfig *config.Config
}

func New(appName, appOwner string, cfgs ...Configurator) *Startup {
	b := &Startup{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func (this *Startup) setupViews() {
	//默认使用后台模板
	this.RegisterView(iris.HTML("./web", ".html").Layout("admin/views/layout/admin_layout.html").Reload(true))
}

func (this *Startup) setupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	this.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + this.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
		AllowReclaim: true,
	})
}


func (this *Startup) setupWebsockets(endpoint string, handler websocket.ConnHandler) {
	ws := websocket.New(websocket.DefaultGorillaUpgrader, handler)
	this.Get(endpoint, websocket.Handler(ws))
}


func (this *Startup) setupErrorHandlers() {
	this.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     this.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("dmin/views/layout/error.html")
	})
}

func (this *Startup) setupAppConfig()  {
	config.LoadConfig("config/config.toml")
	this.AppConfig = config.GetConfig()
}

func (this *Startup)setupXorm()  {
	engine, err := xorm.NewEngine("mysql",this.AppConfig.MySQL.DSN())
	if err!=nil {
		panic(err)
	}

	err = engine.Ping()
	if err != nil {
		panic(err)
	}

	engine.ShowSQL(this.AppConfig.Xorm.ShowSql)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.ShowExecTime(this.AppConfig.Xorm.ShowExecTime)
	engine.SetMapper(core.SameMapper{})

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{},this.AppConfig.Xorm.TablePrefix)
	engine.SetTableMapper(tbMapper)
	engine.Charset("utf8")

	engine.SetMaxIdleConns(this.AppConfig.Xorm.MaxIdleConns)
	engine.SetMaxOpenConns(this.AppConfig.Xorm.MaxOpenConns)
	engine.SetConnMaxLifetime(time.Duration(this.AppConfig.Xorm.MaxLifetime) * time.Second)

    //创建表
	createTable(engine,entitys.User{})
	createTable(engine,entitys.Article{})
	createTable(engine,entitys.Option{})
	createTable(engine,entitys.Tag{})
	createTable(engine,entitys.TagArticle{})

	this.Xorm =engine

	iris.RegisterOnInterrupt(func() {
		engine.Close()
	})
}

func createTable(engine *xorm.Engine,table interface{}) {
	b,_:= engine.IsTableExist(table)
	if !b{
		er := engine.CreateTables(table)
		if er!=nil {
			panic(er)
		}
	}
}


func (this *Startup) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(this)
	}
}

func (this *Startup) Bootstrap() *Startup {
	//设置app配置
	this.setupAppConfig()
	//设置xorm
	this.setupXorm()
	//设置视图
	this.setupViews()
	//设置session
	this.setupSessions(2 * time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	//设置异常页面
	this.setupErrorHandlers()
	//设置静态文件
	this.HandleDir("/static", "./assets")

	this.Use(recover.New())
	this.Use(logger.New())

	return this
}