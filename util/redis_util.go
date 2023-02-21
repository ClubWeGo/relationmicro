package util

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

// const (
// 	url      = "localhost:6379"
// 	password = "123456"
// 	db       = 0 // 选择的db号

// 	// 连接池参数
// 	maxIdle     = 10  // 初始连接数
// 	maxActive   = 10  // 最大连接数
// 	idleTimeOut = 300 // 最长空闲时间
// )

type Config struct {
	Url      string
	Password string
	DB       int // 选择的db号

	// 连接池参数
	MaxIdle     int // 初始连接数
	MaxActive   int // 最大连接数
	IdleTimeOut int // 最长空闲时间
}

var pool *redis.Pool

func Init(config Config) {
	// 加载lua脚本
	InitLoadLua()

	pool = &redis.Pool{
		MaxIdle:     config.MaxIdle,                    //最初的连接数量
		MaxActive:   config.MaxActive,                  //连接池最大连接数量,（0表示自动定义），按需分配
		IdleTimeout: time.Duration(config.IdleTimeOut), //连接关闭时间 300秒 （300秒不使用自动关闭）
		// 连接
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			// 建立tcp连接
			c, err := redis.Dial("tcp", config.Url)
			if err != nil {
				return nil, err
			}
			// 验证密码
			if _, err := c.Do("AUTH", config.Password); err != nil {
				c.Close()
				return nil, err
			}

			// 选择库
			if _, err := c.Do("SELECT", config.DB); err != nil {
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

// zset 从大到小 只返回value
func FindTopVal(key string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("zrevrangebyscore", redis.Args{}.Add(key).Add("+inf").Add("-inf")...))
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

// zset 元素总数量
func FindZSetCount(key string) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("zcard", key))
}

// zset 元素总数量 + 范围
func FindZSetCountByRange(key string, min string, max string) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("zcount", redis.Args{}.Add(key).Add(min).Add(max)...))
}

// ZSet value值是否存在
func FindZSetIsExists(key string, value int64) (bool, error) {
	conn := pool.Get()
	defer conn.Close()
	reply, err := conn.Do("zrank", redis.Args{}.Add(key).Add(value)...)
	if err != nil {
		return false, err
	}
	// 转int 如果结果为nil 会转换失败 则返回false
	if _, err := redis.Int64(reply, nil); err != nil {
		return false, err
	}
	return true, nil

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

/*
*
设置hash的字段值
*/
func HSet(key string, field string, value string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("hset", key, field, value)
}

/*
*
设置hash的字段值 filed int64
*/
func HSetI64(key string, field int64, value string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("hset", key, field, value)
}

/*
*
设置hash的多个字段
*/
func HMSet(key string, fieldAndValues ...string) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	l := len(fieldAndValues)
	var is = make([]interface{}, l, l)
	for i, field := range fieldAndValues {
		is[i] = field
	}
	return redis.Values(conn.Do("hmset", redis.Args{}.Add(key).Add(is...)...))
}

/*
*
获取hash的字段值
*/
func HGet(key string, field string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("hget", key, field))
}

/*
*
获取hash的字段值 field int64
*/
func HGetI64(key string, field int64) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("hget", key, field))
}

/*
*
获取hash的多个字段
*/
func HMGet(key string, fields ...string) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	l := len(fields)
	var is = make([]interface{}, l, l)
	for i, field := range fields {
		is[i] = field
	}
	return redis.Values(conn.Do("hmget", redis.Args{}.Add(key).Add(is...)...))
}

/*
*
获取hash的多个字段 field 为 int64
*/
func HMGetFiledI64(key string, fields ...int64) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	// int64 不能直接添加到 redis参数 需要转换为 []interface{}
	l := len(fields)
	var is = make([]interface{}, l, l)
	for i, field := range fields {
		is[i] = field
	}
	//fmt.Println(redis.Args{}.Add(is...))
	return redis.Values(conn.Do("hmget", redis.Args{}.Add(key).Add(is...)...))
}

/*
*
根据 int64的fields 返回 map<int64, string> field-val
*/
func HMGetI64ReturnMapI64(key string, fields ...int64) (map[int64]string, error) {
	res, err := HMGetFiledI64(key, fields...)
	if err != nil {
		return nil, err
	}
	log.Printf("HMGetFiledI64 end, res:%s", res)
	return ConvertHashFieldI64(fields, res), nil
}

// hmget 返回 需要转换
func ConvertHashFieldI64(fields []int64, resp []interface{}) map[int64]string {
	var resMap = make(map[int64]string)
	idLen := len(fields)
	for i, item := range resp {
		// 避免两集合长度对不上的情况
		if i >= idLen {
			log.Printf("HashConvertFieldI64 根据userIds查出来的userNames长度对不上")
			continue
		}
		var name = "unknown"
		if item != nil {
			name = string(item.([]byte))
		}
		resMap[fields[i]] = name
	}
	return resMap
}

/*
 * message utils
 */

// reference :
// https://blog.csdn.net/li_w_ch/article/details/110638434
// https://juejin.cn/post/7112825943231561741

// 创建一条流消息
func XADD(roomId, msgid, userId, toUserId, value string, maxlen int32) (string, error) {
	// roomId 聊天室房间号
	// msgId 消息号 推荐使用默认 * , 自带时间戳支持范围查询
	// key 标识用户id
	// value 消息内容
	// maxlen 默认-1表示不限制长度，可手动指定大小限制stream的长度
	conn := pool.Get()
	defer conn.Close()
	fromToString := fmt.Sprintf("%v_%v", userId, toUserId)
	if maxlen == -1 {
		return redis.String(conn.Do("XADD", roomId, "*", "from_to", fromToString, "content", value))
	}
	return redis.String(conn.Do("XADD", roomId, "*", "MAXLEN ~", maxlen, "from_to", fromToString, "content", value)) //模糊限制
	// 返回结果为消息的ID, reply -> 1676530008466-0
}

// 删除一条流消息
func XDEL(roomId, msgId string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("XDEL", roomId, msgId))
}

// 删除房间内所有消息
func XDELALL(roomId string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("XTRIM", roomId, "MAXLEN", 0))
}

/*
XREVRANGE 倒序取max-min范围内数据
*/
func XREVRANGE(roomId string, start string, end string) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("XREVRANGE", roomId, start, end))
}

/*
*
eval 执行lua脚本
script：脚本内容
keyNumber：执行脚本需要多少key
keyArgvs key和argv
eval script keyNumber [keys] [argvs]
*/
func Eval(script string, keyNumber int, keyArgvs ...interface{}) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	// 参数转换为 []interface{}
	//is := StrArrToInterfaceArr(keyArgvs)
	// todo 可考虑 redis.NewScript
	return conn.Do("eval", redis.Args{}.Add(script).Add(keyNumber).Add(keyArgvs...)...)
}

/*
*
lua 返回 ints
*/
func EvalReturnInts(script string, keyNumber int, keyArgvs ...interface{}) ([]int, error) {
	return redis.Ints(Eval(script, keyNumber, keyArgvs...))
}

/*
*
根据 lua sha1 直接执行缓存的lua脚本
*/
func EvalSha(luaSha1 string, keyNumber int, keyArgvs ...interface{}) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("EVALSHA", redis.Args{}.Add(luaSha1).Add(keyNumber).Add(keyArgvs...)...)
}

/*
*
根据lua的sha1 查看脚本是否缓存
script exists code1 code2
可以查多个脚本是否存在 所以返回是数组
*/
func ScriptExists(luaSha1 string) ([]int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64s(conn.Do("SCRIPT", "EXISTS", luaSha1))
}

/*
*
redis 返回的数字前面带一个空格 不能直接转
*/
func ScriptExistsRtnInt(luaSha1 string) (int, error) {
	exists, err := ScriptExists(luaSha1)
	if err != nil {
		return 0, err
	}
	return int(exists[0]), nil
}

/*
*
优化版eval
利用eval的脚本缓存机制
*/
func EvalOptimize(script string, keyNumber int, keyArgvs ...interface{}) (interface{}, error) {
	sha1, err := GetLuaSha1(script)
	if err != nil {
		return "", err
	}
	existsCode, err := ScriptExistsRtnInt(sha1)
	if err != nil {
		return "", err
	}
	if existsCode == 0 {
		if _, err := Eval(script, keyNumber, keyArgvs...); err != nil {
			return "", err
		}
	}
	return EvalSha(sha1, keyNumber, keyArgvs...)
}
