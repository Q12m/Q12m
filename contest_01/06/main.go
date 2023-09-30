package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var count_5000 int = n / 5000
	n -= count_5000 * 5000
	var count_1000 int = n / 1000
	n -= count_1000 * 1000
	var count_500 int = n / 500
	n -= count_500 * 500
	var count_200 int = n / 200
	n -= count_200 * 200
	var count_100 int = n / 100
	n -= count_100 * 100
	fmt.Println(count_5000, count_1000, count_500, count_200, count_100)
}
