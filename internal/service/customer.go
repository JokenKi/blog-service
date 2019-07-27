package service

import (
	"crypto/md5"
	"encoding/hex"
	"blog-service/internal/model"
	"strings"

	"github.com/bilibili/kratos/pkg/ecode"

	"math/rand"
	"time"

	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/satori/go.uuid"
)

//校验密码
func CheckPasswd(passwd string, c *model.Customer) bool {
	m5 := md5.New()
	m5.Write([]byte(string(c.Salt)))
	m5.Write([]byte(string(passwd)))
	str := strings.ToUpper(hex.EncodeToString(m5.Sum(nil)))
	log.Info("CheckPasswd input: %s real: %s", str, c.Passwd)
	return strings.Compare(str, c.Passwd) == 0
}

// 获取MD5
func GetMd5(salt int32, passwd string) string {
	m5 := md5.New()
	m5.Write([]byte(string(salt)))
	m5.Write([]byte(string(passwd)))
	return strings.ToUpper(hex.EncodeToString(m5.Sum(nil)))
}

func GetRandNumber(length int) int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
}

//用户登陆
func Login(ctx *bm.Context, s *Service, c *model.Customer) (err error) {
	if len(c.Name) < 1 || len(c.Passwd) < 6 {
		log.Warn("Customer Login got invalid input param: %s : %s", c.Name, c.Passwd)
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	cus := s.dao.GetUserByName(ctx, c.Name)
	if !CheckPasswd(c.Passwd, cus) {
		log.Warn("Customer Login password check failed, name: %s password: ", c.Name, c.Passwd)
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	token := s.dao.GetCustomerToken(ctx, cus.Id)
	if len(token) > 0 {
		ctx.JSON(token, ecode.OK)
		return
	}
	expire := 604800
	token = strings.ToUpper(uuid.Must(uuid.NewV4()).String())
	s.dao.SetCustomerToken(ctx, cus.Id, token, expire)
	s.dao.SetCustomerToCache(ctx, token, cus, expire)
	ctx.JSON(token, ecode.OK)
	return
}

func Regist(ctx *bm.Context, s *Service, c *model.Customer) {
	if len(c.Passwd) < 1 || len(c.Phone) < 1 {
		log.Warn("Customer Regist got invalid input param: %s : %s", c.Phone, c.Passwd)
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	if len(c.Name) < 1 {
		c.Name = c.Phone
	}
	if len(c.NickName) < 1 {
		c.NickName = c.Phone
	}
	userCount := s.dao.CountUser(ctx, c.Name)
	log.Info("Customer  %s userCount %d", c.Name, userCount)
	if userCount > 0 {
		log.Warn("Customer  %s already exist", c.Name)
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	salt := GetRandNumber(6)
	md5Str := GetMd5(salt, c.Passwd)
	nowTimestamp := time.Now().Unix()
	// log.Info("salt  %d md5Str  %s nowTimestamp %d", salt, md5Str, nowTimestamp)

	c.TimeCreate = nowTimestamp
	c.TimeUpdate = nowTimestamp
	c.TimeLatestLogin = 0
	c.Status = 5
	c.Salt = salt
	c.AccountType = 1
	c.Passwd = md5Str
	s.dao.InsertUser(ctx, c)
	ctx.JSON(c.Id, ecode.OK)
}
