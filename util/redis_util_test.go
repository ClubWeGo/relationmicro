package util

import (
	"fmt"
	"testing"
	"time"
)

func TestZadd(t *testing.T) {
	res, err := zadd("k1", time.Now().UnixNano(), 1)
	if err != nil {
		fmt.Printf("zadd err ：%s", err)
	}


	res, err = zrevrangeByScoreOffset("k1", "-inf", "+inf", 0, 10)
	if err != nil {
		return
	}

	fmt.Printf(res.(string))


}


func TestMain(m *testing.M) {
	fmt.Println("begin")
	Init()
	m.Run()
	Close()
	fmt.Println("end")
}
