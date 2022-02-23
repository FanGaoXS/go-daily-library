package main

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

type Student struct {
	Id   int64
	Name string
	Age  int64
	Desc string
}

type Teacher struct {
	Id    int64
	Name  string
	Class Class
}

type Class struct {
	Id   int64
	Name string
}

func randomInt(max int64) int64 {
	if max <= 0 {
		return 0
	}
	// 0 < n < max
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}

func RandomStudent() Student {
	i := randomInt(100)
	return Student{
		Id:   i,
		Name: "user" + strconv.Itoa(int(i)),
		Age:  i * 10,
		Desc: "desc" + strconv.Itoa(int(i)),
	}
}

func RandomTeacher() Teacher {
	i := randomInt(100)
	return Teacher{
		Id:    i,
		Name:  "teacher" + strconv.Itoa(int(i)),
		Class: RandomClass(),
	}
}

func RandomClass() Class {
	i := randomInt(100)
	return Class{
		Id:   i,
		Name: "class" + strconv.Itoa(int(i)),
	}
}
