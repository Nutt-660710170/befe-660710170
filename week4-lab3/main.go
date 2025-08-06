package main

import (
	"fmt"
	"errors"
)

type Student struct {
	ID string `jason:"id"`
	Name string `jason:"name"`
	Email string `jason:"email"`
	Year int `jason:"year"`
	Gpa float64 `jason:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.Gpa >= 3.50
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.Gpa < 0 || s.Gpa > 4 {
		return errors.New("gpa must be between 0-4")
	}
	return nil
}

func main() {
//	var st Student = Student{ID:"1", Name:"nutt", Email:"nutt5346789@gmail.com", Year:3, Gpa:3.75}	
////or--- 	st := Student({ID:"1", Name:"nutt", Email:"nutt5346789@gmail.com", Year:"3", Gpa:"3.01"})
	// fmt.Printf("Honor %v\n",st.IsHonor())
	// fmt.Printf("Validate %v\n",st.Validate())


	students := []Student{
		{ID:"1", Name:"nutt", Email:"nutt5346789@gmail.com", Year:3, Gpa:3.75},
		{ID:"2", Name:"a", Email:"a@gmail.com", Year:2, Gpa:2.75}
	}

	newStudent := Student{ID:"3", Name:"b", Email:"b@gmail.com", Year:1, Gpa:3.50}
	students=append(students,newStudent)

for i, student := range students {
	fmt.Printf("Honor %v\n",i,st.IsHonor())
	fmt.Printf("Validate %v\n",i,st.Validate())
}
	


}