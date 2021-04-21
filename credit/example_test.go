package credit_test // взяли пакет credit, добавили _test

import (
	"fmt"
	"github.com/Fluffi21213/homework/credit"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(37163, 36, 20))
	// Output:
	// 37184 338623 1338623
}
