package util

import (
	"fmt"
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

func TestMain(m *testing.M) {
	//fmt.Println("begin")
	Init()
	m.Run()
	//println(GetFollowedTimeStr())
	Close()
	//fmt.Println("end")
}
