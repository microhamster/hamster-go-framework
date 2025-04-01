package core

import (
	"context"
	"database/sql"
	"hamster/log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

// 获取缓存连接
func OpenRedis(ctx context.Context, endpoint string) *redis.Client {

	options, err := redis.ParseURL(endpoint)
	if err != nil {
		log.Errorf("failed to get redis options: %s", err.Error())
		return nil
	}

	options.PoolSize = 1024
	if !strings.Contains(endpoint, "min_idle_conns") {
		options.MinIdleConns = 8
	}

	client := redis.NewClient(options)
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Errorf("failed to connect redis: %s %s", EVENT_APP_PANIC, err.Error())
		return nil
	}

	return client
}

// 获取数据库连接
func OpenMysql(ctx context.Context, endpoint string) *sql.DB {

	db, err := sql.Open("mysql", endpoint)
	if err != nil {
		log.Errorf("failed to connect mysql: %s %s -> %s", EVENT_APP_PANIC, endpoint, err.Error())
		return nil
	}

	return db
}
