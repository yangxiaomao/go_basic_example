package show

import (
	"fmt"
	"gotest/01/11_package/calc"
)

func Show() {
	a1 := calc.Add(10, 2)
	a2 := calc.Mul(10, 2)
	a3 := calc.Sub(10, 2)
	a4 := calc.Div(10, 2)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}
