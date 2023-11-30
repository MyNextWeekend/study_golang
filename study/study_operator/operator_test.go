package study_operator

import (
	"testing"
)

func TestOperator(t *testing.T) {
	t.Logf("4右移2位是: %v\n", 4>>2)
	t.Logf("2左移1位是: %v\n", 2<<1)
	t.Logf("&按位与运算: %v\n", 2&1) //对应位置二进制相与，即：同为1，结果为1，否则为0
	t.Logf("|按位或运算: %v\n", 2|1) //对应位置二进制相或，即：有一个为1，结果为1，否则为0
}
