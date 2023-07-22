package study

import (
	"testing"
	"unicode/utf8"
)

func TestStr(t *testing.T) {
	nameStr := "阿斯顿法国萨阿发是打发来了"
	t.Logf("len直接计算长度：%v", len(nameStr))
	t.Logf("utf8包计算长度：%v", utf8.RuneCountInString(nameStr))

	nameList := []rune(nameStr)
	t.Logf("转换为数组之后为：%v", nameList)
	t.Logf("转换为数组长度为：%v", len(nameList))
}
