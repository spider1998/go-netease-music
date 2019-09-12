package api

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"music/log"
)

func OpenRedis(addr string, size int, logger log.Logger) (*RedisClient, error) {
	p, err := pool.New("tcp", addr, size)
	if err != nil {
		return nil, err
	}

	c := new(RedisClient)
	c.Pool = p
	c.logger = logger
	return c, nil
}

type RedisClient struct {
	*pool.Pool
	logger log.Logger
}

func (c *RedisClient) Cmd(cmd string, args ...interface{}) *redis.Resp {
	resp := c.Pool.Cmd(cmd, args...)
	c.logger.Debug("redis command.", "cmd", cmd, "args", args, "resp", resp.String())
	return resp
}

type RedisResponse struct {
	*redis.Resp
}

func (r *RedisResponse) IsNil() bool {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.IsType(redis.Nil)
}

func GetRedis(dsn string, logger log.Logger) (Redis *RedisClient, err error) {
	Redis, err = OpenRedis(dsn, 10, logger)
	if err != nil {
		return
	}
	return
}
