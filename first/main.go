package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	str := first(n)
	fmt.Println(str)

}

func first(n int) string {
	if 11 <= n%100 && n%100 <= 14 {
		return fmt.Sprintf("%d компьютеров", n)
	} else if n%10 == 1 {
		return fmt.Sprintf("%d компьютер", n)
	} else if 2 <= n%10 && n%10 <= 4 {
		return fmt.Sprintf("%d компьютера", n)
	} else {
		return fmt.Sprintf("%d компьютеров", n)
	}
}
