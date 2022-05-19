// Package model 数据模型
package example

import (
	"fmt"
	"nature-beta/model"
	"testing"
)

// TestCurl 测试函数
func TestCurl(t *testing.T) {
	c, err := model.ParseTheFile("./curl/test.curl.txt")
	fmt.Println(c, err)
	if err != nil {
		return
	}
	fmt.Printf("curl:%s \n", c.String())
	fmt.Printf("url:%s \n", c.GetURL())
	fmt.Printf("method:%s \n", c.GetMethod())
	fmt.Printf("body:%v \n", c.GetBody())
	fmt.Printf("body string:%v \n", c.GetBody())
	fmt.Printf("headers:%s \n", c.GetHeadersStr())
}
