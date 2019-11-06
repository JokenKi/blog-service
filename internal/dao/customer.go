package dao

import (
	"blog-service/internal/model"
	"context"

	"github.com/bilibili/kratos/pkg/log"
)

const _getUserByNameSQL = "SELECT c.id,c.`name`,c.nick_name,c.passwd,c.salt,c.phone,c.account_type,c.`status`,c.time_create,c.time_update,c.time_latest_login FROM tb_customer c where c.status = 5 and c.name=?"

const _insertUserSQL = "INSERT INTO `tb_customer` (`name`,`nick_name`,`passwd`,`salt`,`phone`,`account_type`,`status`,`time_create`,`time_update`)VALUES(?, ?,?,?,?,?,?,?,?)"

const _countUserByNameSQL = "SELECT COUNT(id) FROM tb_customer WHERE NAME = ?"

const _updatePasswdSQL = "UPDATE tb_customer SET passwd = ?,time_update=? WHERE NAME = ?"

//根据用户名称获取用户信息
func (d *Dao) GetUserByName(ctx context.Context, username string) (cus *model.Customer) {
	cus = &model.Customer{}
	row := d.db.QueryRow(ctx, _getUserByNameSQL, username)
	if err := row.Scan(&cus.Id,
		&cus.Name,
		&cus.NickName,
		&cus.Passwd,
		&cus.Salt,
		&cus.Phone,
		&cus.AccountType,
		&cus.Status,
		&cus.TimeCreate,
		&cus.TimeUpdate,
		&cus.TimeLatestLogin); err != nil {
		cus = nil
		log.Error("row.Scan error(%v)", err)
		return
	}
	return
}

func (d *Dao) CountUser(ctx context.Context, name string) int64 {
	row := d.db.QueryRow(ctx, _countUserByNameSQL, name)
	var count int64
	row.Scan(&count)
	return count
}

//插入用户信息
func (d *Dao) InsertUser(ctx context.Context, cus *model.Customer) int64 {
	stmt, err := d.db.Prepare(_insertUserSQL)
	checkErr(err)
	res, err := stmt.Exec(ctx, &cus.Name, &cus.NickName, &cus.Passwd, &cus.Salt, &cus.Phone, &cus.AccountType, &cus.Status, &cus.TimeCreate, &cus.TimeUpdate)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}

//更新用户密码
func (d *Dao) UpdatePasswd(ctx context.Context, cus *model.Customer) int64 {
	stmt := d.db.Prepared(_updatePasswdSQL)
	res, err := stmt.Exec(ctx, &cus.NewPasswd, &cus.TimeUpdate, &cus.Name)
	checkErr(err)
	rows, err := res.RowsAffected()
	checkErr(err)
	return rows
}
