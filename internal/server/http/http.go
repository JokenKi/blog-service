package http

import (
	"net/http"

	"go-common/library/ecode"
	"blog-service/internal/model"
	"blog-service/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/blog-service")
	{
		g.GET("/start", howToStart)
	}

	c := e.Group("/customer")
	{
		c.GET("/login", login)
		c.POST("/regist", regist)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

//用户登陆
func login(ctx *bm.Context) {
	c := &model.Customer{}
	if err := ctx.Bind(c); err != nil {
		return
	}
	if c.Name == "" || !preLogin(c) {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	service.Login(ctx, svc, c)
}

func regist(ctx *bm.Context) {
	c := &model.Customer{}
	if err := ctx.Bind(c); err != nil {
		return
	}
	service.Regist(ctx, svc, c)
}

func preLogin(c *model.Customer) (invalid bool) {
	invalid = true
	return
}