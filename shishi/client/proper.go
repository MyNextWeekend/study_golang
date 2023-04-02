package client

import "fmt"

type courseClient struct {
}

func NewCourseClient() *courseClient {
	return &courseClient{}
}

func (c *courseClient) LearnGo(ctx int, req string) (res string, err error) {
	fmt.Println("golang学习成功")
	return "", nil
}

func (c *courseClient) LearnJava(ctx int, req string) (res string, err error) {
	fmt.Println("java学习成功")
	return "", nil
}
func (c *courseClient) LearnPython(ctx int, req string) (res string, err error) {
	fmt.Println("python学习成功")
	return "", nil
}
func (c *courseClient) LearnC(ctx int, req string) (res string, err error) {
	return "", nil
}
func (c *courseClient) LearnVUE(ctx int, req string) (res string, err error) {
	return "", nil
}
