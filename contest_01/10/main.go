package main

import "fmt"

func main() {
	var array map[string][3]bool
	array = make(map[string][3]bool)

	array["Кот"] = [3]bool{false, false, false}
	array["Жираф"] = [3]bool{false, false, true}
	array["Курица"] = [3]bool{false, true, false}
	array["Страус"] = [3]bool{false, true, true}
	array["Дельфин"] = [3]bool{true, false, false}
	array["Плезиозавры"] = [3]bool{true, false, true}
	array["Пингвин"] = [3]bool{true, true, false}
	array["Утка"] = [3]bool{true, true, true}

	var ans1, ans2, ans3 string

	fmt.Scan(&ans1, &ans2, &ans3)

	var b_ans1 bool = ans1 == "Да"
	var b_ans2 bool = ans2 == "Да"
	var b_ans3 bool = ans3 == "Да"

	var ar_2 = [3]bool{b_ans1, b_ans2, b_ans3}

	for k, v := range array {
		if ar_2 == v {
			fmt.Println(k)
		}

	}

}

