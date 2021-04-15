package main

import (
	"fmt"
)

func main() {
	a := 3333_33
	b := 1
	limit := 100

	bonus := a * b / 10000 // вместо нуля ваша формула
	if bonus > limit {
		bonus = limit
	}
	fmt.Print(bonus) // должно быть 33* (см. объяснение ниже)
}
