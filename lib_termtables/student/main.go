package main

import (
	"fmt"
	"github.com/apcera/termtables"
)

func main() {
	var studentList []Student
	// generate student list
	for i := 0; i < 10; i++ {
		studentList = append(studentList, RandomStudent())
	}
	if len(studentList) == 0 {
		return
	}
	t := termtables.CreateTable()

	// parse the quick_start element of the list with fields as the header of table
	fields := Fields(studentList[0])
	t.AddHeaders(fields...)
	for _, s := range studentList {
		t.AddRow(Values(s)...)
	}
	fmt.Print(t.Render())

	var teacherList []Teacher
	// generate student list
	for i := 0; i < 10; i++ {
		teacherList = append(teacherList, RandomTeacher())
	}
	if len(teacherList) == 0 {
		return
	}
	t1 := termtables.CreateTable()

	// parse the quick_start element of the list with fields as the header of table
	fields = Fields(teacherList[0])
	t1.AddHeaders(fields...)
	for _, t := range teacherList {
		t1.AddRow(Values(t)...)
	}
	fmt.Print(t1.Render())
}
