package main

import "fmt"

func name(s *string) {
	fmt.Println("666", *s)
}

func main() {
	s := "woshishei"
	fmt.Print(&s)
	name(&s)
}
