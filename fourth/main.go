package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введи желаемый размер таблицы(целое число)")
	fmt.Scan(&n)
	solution(n)
}

func solution(n int) {
	for row := 0; row <= n; row++ {
		for col := 0; col <= n; col++ {
			if row == col && row == 0 {
				fmt.Printf("    ")
			} else if row == 0 {
				fmt.Printf("%4.d", col)
			} else if col == 0 {
				fmt.Printf("%4.d", row)
			} else {
				fmt.Printf("%4.d", row*col)
			}
		}
		fmt.Printf("\n")
	}

}
