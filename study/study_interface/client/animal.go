package client

import "fmt"

type AnimalClient struct {
}

func NewAnimalClient() *AnimalClient {
	return &AnimalClient{}
}

func (c *AnimalClient) LearnGo(ctx int, req string) (res string, err error) {
	fmt.Println("golang学习成功")
	return "", nil
}
func (c *AnimalClient) LearnPython(ctx int, req string) (res string, err error) {
	fmt.Println("python学习成功")
	return "", nil
}
