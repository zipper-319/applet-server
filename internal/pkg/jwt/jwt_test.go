package jwtUtil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetToken(t *testing.T) {
	Convey("测试获取token", t, func() {
		account := "zipper"
		//expire := 100
		key := ""
		tokenStr, err := GetToken(account, "123", 1)
		t.Log(tokenStr)
		So(err, ShouldBeNil)
		tokenInfo, err := ParseToken(tokenStr, key)

		t.Log(tokenInfo)
		So(err, ShouldBeNil)

	})
}
