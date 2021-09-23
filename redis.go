package conngo

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	Addr     string
	Password string
	DB       int

	Opt    *redis.Options //此选项是可选的，如果有此选项，则覆盖其它配置
	client *redis.Client
}

// 获取一个已存在的Client
// 不需要Close
func (r *Redis) GetClient() *redis.Client {
	if r.client == nil {
		r.client = r.NewClient()
	}
	return r.client
}

// 始终获取一个新的redis连接
// 用完需要 Close 释放
func (r *Redis) NewClient() *redis.Client {
	if r.Opt == nil {
		r.Opt = &redis.Options{
			Addr:         r.Addr,
			Password:     r.Password,
			DB:           r.DB,
			MinIdleConns: 6000,
		}
	}
	return redis.NewClient(r.Opt)
}
