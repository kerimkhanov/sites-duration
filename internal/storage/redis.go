package storage

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
)

func InitRedis(config *viper.Viper) (client *redis.Client, err error) {
	redisHost := config.GetString("APP.REDIS.HOST") + ":" + config.GetString("APP.REDIS.PORT")
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Printf("Failed to connect to Redis:", err)
		return nil, err
	}
	return redisClient, nil
}
