package random

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min, max int) int {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
	// 生成并返回随机数
	return rand.Intn(max-min+1) + min
}
