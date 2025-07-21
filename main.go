package main

import "fmt"

func main() {

	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	m := Matrix{}
	m.Data = data
	fmt.Println(getDet(m))

}
