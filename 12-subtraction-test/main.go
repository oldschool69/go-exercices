package main

import "fmt"

func main() {
	avg := Average([]float64{1, 2}...)
	fmt.Println(avg)
	avg = Average([]float64{2, 2, 2}...)
	fmt.Println(avg)
	avg = Average([]float64{1, 2, 3, 4, 5}...)
	fmt.Println(avg)

}


func Average(nums ...float64) float64 {
	var sum float64
	for _,v := range nums {
		sum += v
	}

	return sum / float64(len(nums))
}