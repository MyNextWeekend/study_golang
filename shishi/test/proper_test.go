package test

import (
	"study_golang/shishi/service"
	"testing"
)

func TestAbc(t *testing.T) {
	//client := client.NewCourseClient()
	serv := service.NewCourseService()
	serv.StudyPython()
	serv.StudyGolang()
}
