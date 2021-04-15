package main

import (
	"fmt"
	"math"
)

func main() {
	s := float64(1000000_00)
	i := 0.2 / 12
	n := float64(36)
	k := i * math.Pow(1+i, n) / (math.Pow(1+i, n) - 1)
	a := k * s
	r := a * n
	v := r - s
	y := int(a) / 100
	o := int(r) / 100
	b := int(v) / 100
	fmt.Println(y)
	fmt.Println(b)
	fmt.Println(o)

}
