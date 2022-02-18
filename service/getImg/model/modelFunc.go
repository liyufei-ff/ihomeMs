package model

import "github.com/gomodule/redigo/redis"

var RedisPool redis.Pool

// InitRedis 连接池
func InitRedis() {
	//连接Redis
	RedisPool = redis.Pool{
		MaxIdle:         20,
		MaxActive:       50,
		MaxConnLifetime: 60 * 5,
		IdleTimeout:     60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "47.94.195.58:6379", redis.DialPassword("123456"))
		},
	}
}

// SaveImgRnd 存储验证码
func SaveImgRnd(uuid, rnd string) error {
	//链接redis
	conn := RedisPool.Get()
	//存储验证码
	_, err := conn.Do("setex", uuid, 60*5, rnd)
	return err
}
