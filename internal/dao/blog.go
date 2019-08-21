package dao

import (
	"blog-service/internal/model"
	"context"
)

const _insertBlogSQL = "INSERT INTO `tb_blog` ( `customer_id`, `type_id`, `blog_title`, `content`, `read_num`, `status`, `time_create`, `time_update`) VALUES (?,?,?,?,?,?,?,?)"

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
	return id
}
