package main

import "fmt"

func main() {

	x := [10]float64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 0; i < 10; i++ {

		fmt.Println("Array default value = Inde = ", i, x[i])
	}
}
