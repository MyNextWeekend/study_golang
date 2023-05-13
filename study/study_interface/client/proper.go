package client

import "fmt"

type CourseClient struct {
}

func NewCourseClient() *CourseClient {
	return &CourseClient{}
}

func (c *CourseClient) LearnGo(ctx int, req string) (res string, err error) {
	fmt.Println("golang学习成功")
	return "", nil
}

func (c *CourseClient) LearnJava(ctx int, req string) (res string, err error) {
	fmt.Println("java学习成功")
	return "", nil
}
func (c *CourseClient) LearnPython(ctx int, req string) (res string, err error) {
	fmt.Println("python学习成功")
	return "", nil
}
func (c *CourseClient) LearnC(ctx int, req string) (res string, err error) {
	return "", nil
}
func (c *CourseClient) LearnVUE(ctx int, req string) (res string, err error) {
	return "", nil
}
