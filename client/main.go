package main

import (
	"fmt"
	"hello/server"
)

func main() {
	p := server.Per{}
	(&p).SetName("wo")
	p.SetName("wo123")
	fmt.Println(p.Name())
}
