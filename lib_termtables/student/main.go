package main

import (
	"fmt"
	"github.com/apcera/termtables"
	"strconv"
)

func main() {
	var studentList []Student
	for i := 0; i < 10; i++ {
		s := Student{
			Name: "test" + strconv.Itoa(i),
			Age:  10 + i,
		}
		studentList = append(studentList, s)
	}
	if len(studentList) == 0 {
		return
	}
	t := termtables.CreateTable()
	t.AddHeaders(Fields(studentList[0]))
	for _, s := range studentList {
		t.AddRow(Values(s))
	}
	fmt.Print(t.Render())
}
