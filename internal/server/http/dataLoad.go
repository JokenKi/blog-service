package http

import (
	"blog-service/internal/service"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func initDataLoadRouter(e *bm.Engine) {
	g := e.Group("/dataLoad")
	{
		g.GET("/loadUserBlogs", loadUserBlogs)
	}
}

func loadUserBlogs(ctx *bm.Context) {
	service.LoadUserBlogs(ctx, svc)
}
