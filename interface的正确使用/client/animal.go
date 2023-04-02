package client

import "fmt"

type animalClient struct {
}

func NewAnimalClient() *animalClient {
	return &animalClient{}
}

func (c *animalClient) LearnGo(ctx int, req string) (res string, err error) {
	fmt.Println("golang学习成功")
	return "", nil
}
func (c *animalClient) LearnPython(ctx int, req string) (res string, err error) {
	fmt.Println("python学习成功")
	return "", nil
}
