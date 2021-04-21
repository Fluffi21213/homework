package credit_test // взяли пакет credit, добавили _test

import (
	"fmt"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(homework.Calculate(1_000_000_00, 36, 20))
	fmt.Println(homework.Calculate(10_000_00, 36, 20))
	// Output:
	// 3718400 33862300 133862300
	// 37184 338623 1338623
}
