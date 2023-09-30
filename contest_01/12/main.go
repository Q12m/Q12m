package main

import "fmt"

func main() {

	var n, step int
	fmt.Scan(&n)

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		step += 1

	}
	fmt.Println(step)

}
