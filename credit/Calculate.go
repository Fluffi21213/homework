package credit

import (
	"math"
)

func Calculate(o, l, y int) (int, int, int) {
	s := float64(1000000_00)
	i := 0.2 / 12
	n := float64(36)
	k := i * math.Pow(1+i, n) / (math.Pow(1+i, n) - 1)
	a := k * s
	r := a * n
	v := r - s
	o = int(a) / 100
	l = int(v) / 100
	y = int(r) / 100
	return o, y, l

}
