package main

import "fmt"

func main() {
	first(6)

}

func first(n int) {
	if 11 <= n%100 && n%100 <= 14 {
		fmt.Printf("%d компьютеров", n)
	} else if n%10 == 1 {
		fmt.Printf("%d компьютер", n)
	} else if 2 <= n%10 && n%10 <= 4 {
		fmt.Printf("%d компьютера", n)
	} else {
		fmt.Printf("%d компьютеров", n)
	}
}
