package util

import (
	"fmt"
	"log"
	"testing"
)

var testKey = "testKey1"
var testVal = 1

func TestZadd(t *testing.T) {
	res, err := Zadd(testKey, GetFollowedTimeStr(), int64(testVal))
	if err != nil {
		fmt.Printf("zadd err ：%s", err)
	}

	fmt.Println(res, "zadd success")

	ans, err := ZrevrangeByScoreOffset("k1", "-inf", "+inf", 0, 10)
	if err != nil {
		fmt.Println("zrevrange err :", err)
		return
	}

	for _, v := range ans {
		fmt.Printf("%s\n", v.([]byte))
	}

}

func TestZrem(t *testing.T) {
	//zrem()
}

func TestWithScoreConvert(t *testing.T) {
	followList, err := FindTop("relation_follow_1")
	if err != nil {
		fmt.Println(err)
	}
	res := WithScoreConvert(followList)
	for k, v := range res {
		fmt.Println(k, v)
	}
}

func TestHMGetI64ReturnMapI64(t *testing.T) {
	key := GetUserNameKey()
	userIds := []int64{1, 2, 3, 4}
	resMap, err := HMGetI64ReturnMapI64(key, userIds...)
	if err != nil {
		log.Fatal("TestHMGetI64ReturnMapI64 exception：", err)
	}
	for k, v := range resMap {
		fmt.Println(k, v)
	}

}

func TestFindZSetCount(t *testing.T) {
	key := GetFollowKey(12)

	count, err := FindZSetCount(key)
	if err != nil {
		log.Println("TestFindZSetCount exception:", err)
	}
	fmt.Println("zsetCount:", count)
}

// zset 存在
func TestFindZSetExists(t *testing.T) {
	key := GetFollowKey(12)
	exists, err := FindZSetIsExists(key, 2)
	if err != nil {
		log.Fatal("TestFindZSetExists exception:", err)
	}
	fmt.Printf("TestFindZSetExists exists: %v", exists)
}

// zset 不存在
func TestFindZSetNoneExists(t *testing.T) {
	key := GetFollowKey(12)
	exists, err := FindZSetIsExists(key, 0)
	if err != nil {
		log.Fatal("TestFindZSetNoneExists exception:", err)
	}
	fmt.Printf("TestFindZSetNoneExists exists: %v", exists)
}

func TestHSetI64(t *testing.T) {
	key := GetUserNameKey()
	if _, err := HSetI64(key, 3, "zhang3"); err != nil {
		t.Errorf("TestHSetI64 exception: %s", err)
	}
}

func TestHGetI64(t *testing.T) {
	key := GetUserNameKey()
	name, err := HGetI64(key, 3)
	if err != nil {
		t.Errorf("TestHGetI64 exception: %s", err)
	}
	fmt.Println(name)
}

func TestMain(m *testing.M) {
	//fmt.Println("begin")

	config := Config{
		Url:         "localhost:6379",
		Password:    "123456",
		DB:          0,
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeOut: 300,
	}

	Init(config)
	m.Run()
	//println(GetFollowedTimeStr())
	Close()
	//fmt.Println("end")
}
