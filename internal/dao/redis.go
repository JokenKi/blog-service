package dao

import (
	"context"
	"encoding/json"

	"go-common/library/cache/redis"
	"blog-service/internal/model"

	"fmt"

	"github.com/bilibili/kratos/pkg/log"
)

const _customerTokenKey = "blog:customer:token:%d"
const _customerInfoKey = "blog:customer:info:%s"

func customerInfoTokenKey(uid int64) string {
	return fmt.Sprintf(_customerTokenKey, uid)
}

func customerInfoKey(token string) string {
	return fmt.Sprintf(_customerInfoKey, token)
}

func (d *Dao) GetCustomerToken(ctx context.Context, uid int64) string {
	if uid < 1 {
		return ""
	}
	conn := d.redis.Get(ctx)
	defer conn.Close()

	tokenKey := customerInfoTokenKey(uid)

	token, err := redis.String(conn.Do("GET", tokenKey))
	if err != nil {
		log.Error("GetCustomerToken error ,key: %s", tokenKey)
		return ""
	}
	log.Info("GetCustomerToken: %v", token)
	return token
}

func (d *Dao) SetCustomerToken(ctx context.Context, uid int64, token string, expire int) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	key := customerInfoTokenKey(uid)

	_, err = conn.Do("SETEX", key, expire, token)
	if err != nil {
		log.Error("conn.Do(SETEX, %s) error(%v)", key, err)
	}
	return
}

func (d *Dao) GetCustomerInfoByToken(ctx context.Context, token string) (cus *model.Customer, err error) {
	if len(token) < 0 {
		return
	}
	conn := d.redis.Get(ctx)
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", customerInfoKey(token)))
	log.Info("GetCustomerInfoByToken: %v", value)
	cus = &model.Customer{}
	json.Unmarshal([]byte(value), &cus)
	return
}

func (d *Dao) SetCustomerToCache(ctx context.Context, tokenStr string, cus *model.Customer, expire int) (err error) {
	jsonCus, error := json.Marshal(cus)
	if error == nil {
		log.Info("SetCustomerToCache: %v", string(jsonCus))
	}
	key := customerInfoKey(tokenStr)

	log.Info("tokenStr: %v key: %v", tokenStr, key)

	conn := d.redis.Get(ctx)
	defer conn.Close()
	_, err = conn.Do("SETEX", key, expire, string(jsonCus))
	if err != nil {
		log.Error("conn.Do(SETEX, %s) error(%v)", key, err)
	}
	return
}
