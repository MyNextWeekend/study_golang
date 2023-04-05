package test

import (
	"study_golang/interface_test/service"
	"testing"
)

func TestAbc(t *testing.T) {
	//client := client.NewCourseClient()
	serv := service.NewCourseService()
	serv.StudyPython()
	serv.StudyGolang()
}
