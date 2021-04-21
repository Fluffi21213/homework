package credit // взяли пакет credit, добавили _test

import (
	"fmt"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(Calculate(37163, 1337889, 337889))
	fmt.Println(Calculate(37163, 1337889, 337889))

	// Output:
	// 37163 1337889 337889
	// 37163 1337889 337889
}
