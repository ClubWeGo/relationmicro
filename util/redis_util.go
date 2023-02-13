package util

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

const (
	url = "localhost:6379"
	password = "123456"
	db = 1 // 选择的db号

	// 连接池参数
	maxIdle = 10         // 初始连接数
	maxActive = 10       // 最大连接数
	idleTimeOut = 300    // 最长空闲时间
)

var pool *redis.Pool


func Init() {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,  //最初的连接数量
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
		fmt.Printf("PING err %s", err)
	}

	fmt.Println("redis pool init success")

}

func Close()  {
	if err := pool.Close(); err != nil {
		fmt.Printf("close redis pool error ： %s", err)
	}
}

func zadd(key string, score int64, value int64) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("zadd", key, score, value)
}

func zrem(key string, value int64) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("zrem", key, value)
}


/**
倒序取范围内数据
 */
func zrevrangeByScore(key string, min string, max string) (interface{}, error){
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("zrevrangebyscore", key, min, max)
}

func zrevrangeByScoreOffset(key string, min string, max string, offset int, limit int) (interface{}, error){
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("zrevrangebyscore", key, min, max, offset, limit)
}