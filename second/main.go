package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите длину массива")
	fmt.Scanln(&n)
	array := solution(n)
	fmt.Println(array)

}

func solution(n int) []int {
	var array = makeArray(n)
	var all_number [][]int
	for _, elem := range array {
		all_number = append(all_number, findDel(elem))
	}
	var counter map[int]int = make(map[int]int, 1000)
	var result_array []int
	for _, elem := range all_number {
		for _, el := range elem {
			counter[el]++
			if counter[el] == len(array) {
				result_array = append(result_array, el)
			}
		}
	}
	return result_array
}
func findDel(n int) []int {
	d := 2
	var array []int
	for d*d < n {
		if n%d == 0 {
			array = append(array, n/d)
			array = append(array, d)
		}
		d += 1
	}
	if d*d == n {
		array = append(array, d)
	}
	return array
}

func makeArray(n int) []int {
	var array []int

	for i := 0; i < n; i++ {
		var number int
		fmt.Printf("Введите %d элемент массива\n", i)
		fmt.Scan(&number)
		array = append(array, number)
	}
	return array
}
