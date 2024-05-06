package main

import "fmt"

func main() {
	fmt.Println("Введите левую и правую границу интервала")
	var left, right int
	fmt.Scan(&left)
	fmt.Scan(&right)
	array := third(left, right)
	fmt.Println(array)
}

func third(left, right int) []int {
	var array []int
	for i := left; i <= right; i++ {
		if len(findDel(i)) == 0 {
			array = append(array, i)
		}
	}
	return array
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
