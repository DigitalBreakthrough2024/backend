package database

import (
	"backend/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/spf13/viper"
)

type Session interface {
	Set(ctx context.Context, key string, data SessionData) error
}

type SessionData struct {
	Vector  []float32 `json:"vector"`
	Watched []string  `json:"watched"`
}

type RedisSession struct {
	rdb               *redis.Client
	sessionExpiration time.Duration
	dbResponseTime    time.Duration
}

func InitRedisSession() Session {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString(config.SessionHost),
			viper.GetInt(config.SessionPort),
		),
		Password: viper.GetString(config.SessionPassword),
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to redis: %s", err.Error()))
	}

	return RedisSession{
		rdb:               rdb,
		sessionExpiration: time.Duration(viper.GetInt(config.SessionSaveTime)) * time.Hour * 24,
		dbResponseTime:    time.Duration(viper.GetInt(config.DBResponseTime)) * time.Second,
	}
}

func (r RedisSession) Set(ctx context.Context, key string, data SessionData) error {
	ctx, cancel := context.WithTimeout(ctx, r.dbResponseTime)
	defer cancel()

	sessionDataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = r.rdb.Set(ctx, key, sessionDataJSON, r.sessionExpiration).Err()
	if err != nil {
		return err
	}

	return nil
}
