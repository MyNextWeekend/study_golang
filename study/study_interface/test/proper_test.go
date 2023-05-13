package test

import (
	"study_golang/study/study_interface/service"
	"testing"
)

func TestAbc(t *testing.T) {
	//client := client.NewCourseClient()
	serv := service.NewCourseService()
	serv.StudyPython()
	serv.StudyGolang()
}
