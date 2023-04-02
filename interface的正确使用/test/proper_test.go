package test

import (
	"study_golang/interface的正确使用/service"
	"testing"
)

func TestAbc(t *testing.T) {
	//client := client.NewCourseClient()
	serv := service.NewCourseService()
	serv.StudyPython()
	serv.StudyGolang()
}
