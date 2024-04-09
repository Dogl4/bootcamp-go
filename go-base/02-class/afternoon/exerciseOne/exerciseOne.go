package exercise

import (
	"fmt"
	"time"
)

func ExerciseOne(newStudentsInput ...Student) {
	fmt.Println("Exercício 1:")

	var newStudents = students{}

	newStudents = newStudents.setStudent(newStudentsInput...)

	for _, currentStudent := range newStudents {
		currentStudent.printStudent()
	}
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

type Student struct {
	Name          string
	LastName      string
	Rg            int
	AdmissionDate Date
}

type students []Student

func (s *Student) printStudent() {
	fmt.Printf("Nome: %v\n", s.Name)
	fmt.Printf("Sobrenome: %v\n", s.LastName)
	fmt.Printf("RG: %v\n", s.Rg)
	fmt.Printf("Data de admissão: %v/%v/%v\n", s.AdmissionDate.Day, s.AdmissionDate.Month, s.AdmissionDate.Year)
}

func (ss students) setStudent(student ...Student) students {
	return append(ss, student...)
}
