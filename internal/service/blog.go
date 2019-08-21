package service

import (
	"blog-service/internal/model"

	"github.com/bilibili/kratos/pkg/ecode"

	"time"

	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

//用户发博客
func PublishBlog(ctx *bm.Context, s *Service, c *model.Blog) {
	if c.CustomerId < 1 || c.TypeId < 1 || len(c.BlogTitle) < 1 || len(c.Content) < 1 {
		log.Warn("PublishBlog got invalid input param: %d : %d : %s : %s", c.CustomerId, c.TypeId, c.BlogTitle, c.Content)
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	nowTimestamp := time.Now().Unix()
	// log.Info("salt  %d md5Str  %s nowTimestamp %d", salt, md5Str, nowTimestamp)

	c.TimeCreate = nowTimestamp
	c.TimeUpdate = nowTimestamp
	c.Status = 5
	c.ReadNum = 0
	s.dao.InsertBlog(ctx, c)
	ctx.JSON(c.Id, ecode.OK)
}
