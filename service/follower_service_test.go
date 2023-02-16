package service

import (
	"fmt"
	"testing"
)

func TestFindFollowerList(t *testing.T) {
	list, err := FindFollowerList(1, 2)
	if err != nil {
		t.Errorf("%s", err)
	}
	for _, u := range list {
		fmt.Println(u)
	}
}
