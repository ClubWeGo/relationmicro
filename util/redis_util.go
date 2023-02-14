package util

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

const (
	url      = "localhost:6379"
	password = "123456"
	db       = 0 // 选择的db号

	// 连接池参数
	maxIdle     = 10  // 初始连接数
	maxActive   = 10  // 最大连接数
	idleTimeOut = 300 // 最长空闲时间
)

var pool *redis.Pool

func Init() {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最初的连接数量
		MaxActive:   maxActive,   //连接池最大连接数量,（0表示自动定义），按需分配
		IdleTimeout: idleTimeOut, //连接关闭时间 300秒 （300秒不使用自动关闭）
		// 连接
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			// 建立tcp连接
			c, err := redis.Dial("tcp", url)
			if err != nil {
				return nil, err
			}
			// 验证密码
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}

			// 选择库
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	conn := pool.Get() //从连接池，取一个链接
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		log.Printf("PING err %s", err)
	}

	log.Println("redis pool init success")

}

func Close() {
	if err := pool.Close(); err != nil {
		fmt.Printf("close redis pool error ： %s", err)
	}
}

func Zadd(key string, score string, value int64) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("zadd", key, score, value)
}

// zset 删除 key
func Zrem(key string, value int64) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("zrem", key, value)
}

/*
*
zset 倒序取max-min范围内数据
*/
func ZrevrangeByScore(key string, min string, max string) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("zrevrangebyscore", key, max, min, "withscores"))
}

// zset 倒序取范围内 max-min 的 数据 + 偏移
func ZrevrangeByScoreOffset(key string, min string, max string, offset int, limit int) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("zrevrangebyscore", key, max, min, "withscores", "limit", offset, limit))
}

// zset 高水位 从大到小
func FindTop(key string) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("zrevrangebyscore", key, "+inf", "-inf", "withscores"))
}

// zset 从大到小 + 偏移
func FindTopOffset(key string, offset int, limit int) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("zrevrangebyscore", key, "+inf", "-inf", "withscores", "limit", offset, limit))
}

// zset 从小到大 + 偏移
func FindLow(key string) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("zrevrangebyscore", key, "+inf", "-inf", "withscores"))
}

// zset  从小到大 + 偏移
func FindLowOffset(key string, offset int, limit int) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("zrevrangebyscore", key, "+inf", "-inf", "withscores", "limit", offset, limit))
}

// withscore 返回 需要转换
func WithScoreConvert(resp []interface{}) map[string]string {
	var res = make(map[string]string)
	var key, score = "", ""
	for i, v := range resp {
		if i%2 == 0 {
			//json.Unmarshal(v.([]byte), &item.val)
			// todo 不知道有字符集乱码情况没有 目前没发现
			key = string(v.([]byte))
		} else {
			//json.Unmarshal(v.([]byte), &item.score)
			score = string(v.([]byte))
			res[key] = score
		}
	}
	return res
}
