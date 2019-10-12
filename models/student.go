package models

type Student struct {
	Lastname  string
	Firstname string
	Age       int
}

var StudentList = []Student{
	Student{Lastname: "1", Firstname: "1", Age: 1},
	Student{Lastname: "2", Firstname: "2", Age: 2},
	Student{Lastname: "3", Firstname: "3", Age: 3},
}
