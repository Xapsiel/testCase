package main

import "fmt"

func main() {
	solution()
}

func solution() {
	n :=
	h := 2
	step := 2
	flag := false
	count := 0
	for {
		if n == h {
			break
		}
		count++
		if n <= h+step || flag {
			flag = true
			h += 2
		} else {
			h += step
			step += 1
		}
	}
	fmt.Println(h, count)
}
