package main

import "fmt"
import "github.com/golang-collections/collections/stack"

func main() {

	testExps := []string{
		"10 * ((2 + 5) - 1) * 10 - (3)",
		"1 + (23 + 4)) + (12 / 8)",
		"1 + 2 + 3",
		"10 + (9*(4 - (3 * 5)) + 45)",
		"(((())))",
		"((((3)))",
		"(2(3(3(1, 9, 2)4)4)4)",
		"((((((((((((",
		")))))))))))))))))))(",
		"1*(-10(8*3^2)) + 90 * ((34 + 2) / 23) - 10",
	}

	for _, v := range testExps {
		status := "invalid"
		if isExpValid(v) {
			status = "valid"
		}
		fmt.Printf("exp: %v is %v\n", v, status)
	}

	fmt.Println("-----------------")
	for _, v := range testExps {
		status := "invalid"
		if isExpValid2(v) {
			status = "valid"
		}
		fmt.Printf("exp: %v is %v\n", v, status)
	}

}

// Version using map
func isExpValid(exp string) bool {
	m := make(map[int32]int)
	for _, v := range exp {
		if v == '(' || v == ')'  {
			m[v]++
		}
	}
	return m['('] == m[')']
}

// Version using stack
func isExpValid2(exp string) bool {

	s := stack.New()
	for _, v := range exp {
		if v == '(' {
			s.Push(v)
		} else if v == ')' {
			if s.Len() == 0 {
				s.Push(v)
			} else {
				s.Pop()
			}
		}
	}

	return s.Len() == 0
}
