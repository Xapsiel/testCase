package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите длину массива")
	fmt.Scanln(&n)
	var array = makeArray(n)
	var all_number [][]int
	for _, elem := range array {
		all_number = append(all_number, findDel(elem))
	}
	var result map[int][]int = make(map[int][]int, 1000)
	var result_array []int
	for _, elem := range all_number {
		for _, el := range elem {
			result[el] = append(result[el], 1)
			if len(result[el]) == len(array) {
				result_array = append(result_array, el)
			}
		}
	}
	fmt.Println(result_array)

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
		fmt.Scan(&number)
		array = append(array, number)
	}
	return array
}
