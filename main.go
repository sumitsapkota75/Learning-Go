package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	fmt.Println("AGE:", s.age)
	return s.age
}

func (s *Student) setAge(age int) int {
	fmt.Println("age from setAge 1", s.age)
	s.age = age
	fmt.Println("age from setAge", s.age)
	return s.age
}
func (s Student) getAverageGrage() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	s1 := &Student{name: "Sumit", grades: []int{1, 2, 3}, age: 21}
	s1.getAge()
	s1.setAge(44)
	fmt.Println("Final age", s1.age)
	avg := s1.getAverageGrage()
	fmt.Println("Average", avg)
}
