package conngo

import "github.com/go-redis/redis"

func InitRedis(addr, password string, db int) *Redis {
	return &Redis{addr: addr, password: password, db: db}
}

func InitRedisOpt(opt *redis.Options) *Redis {
	return &Redis{opt: opt}
}

func InitEs(url, username, password string) *Es {
	return &Es{url: url, username: username, password: password}
}

func InitMq(username, password, hostname, endpoint string) *Mq {
	return &Mq{username: username, password: password, hostname: hostname, endpoint: endpoint}
}

func InitMySql(connStr string) *MySql {
	return &MySql{connStr: connStr}
}
