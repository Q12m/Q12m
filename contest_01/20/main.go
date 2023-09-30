package main

import "fmt"

func main() {
	var fst, sec int
	fmt.Scan(&fst, &sec)
	var f, s []int
	f = make([]int, 10)
	s = make([]int, 10)
	for fst > 0 {

		f[fst%10] += 1
		fst /= 10
	}
	for sec > 0 {

		s[sec%10] += 1
		sec /= 10
	}

	flag := true
	for k, v := range f {
		if v != s[k] {
			flag = false
			fmt.Println("NO")
			break
		}
	}
	if flag {
		fmt.Println("YES")
	}

}
