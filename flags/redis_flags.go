package flags

import (
	"github.com/go-redis/redis/v8"
)

type RedisFlags struct {
	Host     string `kong:"required,group=Redis,name=redis-host,env=REDIS_HOST,help=localhost:6379"`
	Username string `kong:"optional,group=Redis,name=redis-username,env=REDIS_USERNAME,default=''"`
	Password string `kong:"optional,group=Redis,name=redis-password,env=REDIS_PASSWORD,default=''"`
}

func (f RedisFlags) Init() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     f.Host,
		Password: f.Password,
		Username: f.Username,
		DB:       0,
	})
}
