package main

import (
	"fmt"
	"plugin"
	"testing"
	"time"
)

func TestEncV1Sign(t *testing.T) {
	// 加载.so插件文件
	p, err := plugin.Open("./enc.so")
	if err != nil {
		t.Fatal("Error loading plugin:", err)
	}

	// 查找插件中的EncV1Sign函数
	encV1SignSym, err := p.Lookup("EncV1Sign")
	if err != nil {
		t.Fatal("Error finding symbol 'EncV1Sign':", err)
	}

	// 类型断言
	encV1SignFunc, ok := encV1SignSym.(func([]byte, []byte, int64, string, int, []byte) ([]byte, error))
	if !ok {
		t.Fatal("Error asserting function type")
	}

	// 定义测试数据
	key := []byte("your_key_here")
	header := []byte("your_header_here")
	timestamp := time.Now().Unix()
	ID := "your_ID_here"
	txType := 0
	body := []byte("your_body_here")

	// 调用函数
	result, err := encV1SignFunc(key, header, timestamp, ID, txType, body)
	if err != nil {
		t.Fatal("Error calling EncV1Sign:", err)
	}

	// 进行其他测试，比如检查结果是否为预期等
	fmt.Println("EncV1Sign returned:", result)
}
