package service

import (
	"fmt"
	"study_golang/interface的正确使用/client"
)

/*
将接口声明延迟到实现类中。
这里只使用到了client中的两个方法，这里只申明两个方法
*/
type clientProxy interface {
	LearnPython(ctx int, req string) (res string, err error)
	LearnGo(ctx int, req string) (res string, err error)
}

type CourseService struct {
	CourseClient clientProxy
}

//// 方案一：通过参数传进来
//func NewCourseService(courseClient clientProxy) *CourseService {
//	return &CourseService{
//		CourseClient: courseClient,
//	}
//}

// NewCourseService 方案二：方法中固定
func NewCourseService() *CourseService {
	return &CourseService{
		CourseClient: client.NewCourseClient(),
	}
}
func (s *CourseService) StudyPython() {
	python, err := s.CourseClient.LearnPython(1, "aaa")
	if err != nil {
		return
	}
	fmt.Print(python)
}

func (s *CourseService) StudyGolang() {
	python, err := s.CourseClient.LearnGo(1, "aaa")
	if err != nil {
		return
	}
	fmt.Print(python)
}
