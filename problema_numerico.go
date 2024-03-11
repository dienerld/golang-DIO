package main

import "fmt"

func main() {
	// append only %3 == 0
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
