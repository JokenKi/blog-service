package http

import (
	"blog-service/internal/service"

	"strconv"

	"github.com/bilibili/kratos/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func initApiRouter(e *bm.Engine) {
	g := e.Group("/api")
	{
		g.GET("/getUserBlogs", getUserBlogs)
	}
}

func getUserBlogs(ctx *bm.Context) {
	userId := ctx.Request.FormValue("userId")
	pageNum := ctx.Request.FormValue("pageNum")
	pageSize := ctx.Request.FormValue("pageSize")
	if len(userId) < 1 {
		paramError(ctx)
		return
	}
	// string到int64
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.JSON("参数错误", ecode.RequestErr)
		return
	}
	pageNumInt := 1
	pageSizeInt := 10
	if len(pageNum) > 1 {
		// string到int
		pageNumInt, err = strconv.Atoi(pageNum)
		if err != nil {
			ctx.JSON("参数错误", ecode.RequestErr)
			return
		}
	}
	if len(pageSize) > 1 {
		// string到int
		pageSizeInt, err = strconv.Atoi(pageSize)
		if err != nil {
			ctx.JSON("参数错误", ecode.RequestErr)
			return
		}
	}

	result := service.GetUserBlogs(ctx, svc, userIdInt64, pageNumInt, pageSizeInt)
	ctx.JSONMap(result, ecode.OK)
}
