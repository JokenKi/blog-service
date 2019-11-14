package dao

import (
	"blog-service/internal/model"
	"container/list"
	"context"

	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/log"
)

const _insertBlogSQL = "INSERT INTO `tb_blog` ( `customer_id`, `type_id`, `blog_title`, `content`, `read_num`, `status`, `time_create`, `time_update`) VALUES (?,?,?,?,?,?,?,?)"
const _BPS_BLOG_USER_LIST_ZSET = "BPS:BLOG:USER:LIST:ZSET."
const _selectBlogSQL = "SELECT tb_blog.id, tb_blog.customer_id, tb_blog.type_id, tb_blog.blog_title, tb_blog.content, tb_blog.read_num, tb_blog.`status`, tb_blog.time_create, tb_blog.time_update FROM tb_blog WHERE tb_blog.`status` = 5"

func (d *Dao) InsertBlog(ctx context.Context, blog *model.Blog) int64 {
	stmt, err := d.db.Prepare(_insertBlogSQL)
	checkErr(err)
	res, err := stmt.Exec(ctx,
		&blog.CustomerId,
		&blog.TypeId,
		&blog.BlogTitle,
		&blog.Content,
		&blog.ReadNum,
		&blog.Status,
		&blog.TimeCreate,
		&blog.TimeUpdate)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	d.AddToCache(ctx, blog, id)
	return id
}

func (d *Dao) AddToCache(ctx context.Context, blog *model.Blog, userId int64) {
	jsonBlog, err := json.Marshal(blog)
	checkErr(err)
	conn := d.redis.Get(ctx)
	defer conn.Close()
	_, err = conn.Do("ZADD", _BPS_BLOG_USER_LIST_ZSET+strconv.FormatInt(int64(userId), 10), blog.TimeUpdate, string(jsonBlog))
	checkErr(err)
}

func (d *Dao) SelectAllBlogs(c context.Context) (blogs *list.List) {
	querySQL := _selectBlogSQL
	rows, err := d.db.Query(c, querySQL)
	if err != nil {
		log.Errorv(c, log.KV("event", "mysql_query"), log.KV("error", err), log.KV("sql", querySQL))
		return
	}
	defer rows.Close()

	for rows.Next() {
		blog := new(model.Blog)
		if err = rows.Scan(&blog.Id,
			&blog.CustomerId,
			&blog.TypeId,
			&blog.BlogTitle,
			&blog.Content,
			&blog.ReadNum,
			&blog.Status,
			&blog.TimeCreate,
			&blog.TimeUpdate); err != nil {
			log.Errorv(c, log.KV("event", "mysql_scan"), log.KV("error", err), log.KV("sql", querySQL))
			return
		}
		if blogs == nil {
			blogs = list.New()
		}
		blogs.PushBack(blog)
	}
	log.Infov(c, log.KV("event", "mysql_query"), log.KV("row_num", blogs.Len()), log.KV("sql", querySQL))
	return
}

func (d *Dao) SelectAllBlogsFromCache(ctx context.Context, userId int64,
	pageNum int,
	pageSize int) (blogs *list.List) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	// 读取指定zset
	user_map, err := conn.Do("ZREVRANGE", _BPS_BLOG_USER_LIST_ZSET+strconv.FormatInt(int64(userId), 10), (pageNum-1)*pageSize, pageSize-1, "WITHSCORES")
	if err != nil {
		fmt.Println("redis get failed:", err)
		return nil
	} else {
		result, err := redis.StringMap(user_map, nil)
		if err != nil {
			panic(err)
		}
		for r := range result {
			blog := new(model.Blog)
			fmt.Printf("user name: %v %v\n", r, result[r])
			json.Unmarshal([]byte(r), blog)
			if blogs == nil {
				blogs = list.New()
			}
			log.Info("id %v", blog.Id)
			blogs.PushBack(blog)
		}
		return blogs
	}
}
