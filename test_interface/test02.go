package main

import "fmt"

type raper interface {
	rap()
}
type downer interface {
	down()
}

type feifei struct {
}

func (f *feifei) down() {
	fmt.Println("下载")
}
func (f *feifei) rap() {
	fmt.Println("唱跳rap")
}

func main() {
	var r raper
	var d downer
	var f feifei
	f = feifei{}
	f.rap()
	f.down()
	d = &f
	d.down()
	r = &f
	r.rap()

}
