package main

import (
	"errors"
	"fmt"
	"strings"
)

// 定义全局变量num
var num int64 = 10

type calculation func(int, int) int

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	sum := 0
	for _, v := range users {
		for _, vv := range v {
			switch vv {
			case 'e', 'E':
				distribution[v] += 1
				sum += 1
			case 'i', 'I':
				distribution[v] += 2
				sum += 2
			case 'o', 'O':
				distribution[v] += 3
				sum += 3
			case 'u', 'U':
				distribution[v] += 4
				sum += 4
			default:
				distribution[v] += 0
				sum += 0
			}
		}

	}
	// fmt.Println(sum)
	// fmt.Println(distribution)
	left = coins - sum
	return left
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	// a := intSum(10, 20)
	// fmt.Println(a)
	// sayHello()
	// a := intSum2(1, 2, 3, 4, 5, 6)
	// fmt.Println(a)
	// ret1 := intSum2()
	// ret2 := intSum2(10)
	// ret3 := intSum2(10, 20)
	// ret4 := intSum2(10, 20, 30)
	// fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60
	// ret5 := intSum3(100)
	// ret6 := intSum3(100, 10)
	// ret7 := intSum3(100, 10, 20)
	// ret8 := intSum3(100, 10, 20, 30)
	// fmt.Println(ret5, ret6, ret7, ret8) // 100 110 130 160
	// a, b := calc(1, 2)
	// fmt.Println(a, b)
	// fmt.Printf("num=%d\n", num) //函数中可以访问全局变量的num
	// testLocalVar()
	// fmt.Println(x) // 此时无法使用变量x
	// testNum() // num = 100
	// testLocalVar2(10, 30)
	// var c calculation
	// c = add
	// fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	// a := c(10, 5)
	// f := add
	// fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	// fmt.Println(f(10, 20))

	// fmt.Println(a)
	// ret2 := gcalc(10, 20, gadd)
	// fmt.Println(ret2) //30
	// a, _ := do("+")
	// fmt.Printf("%T", a)

	// 将匿名函数保存到变量中
	// add := func(x, y int) {
	// 	fmt.Println(x + y)
	// }
	// add(10, 20) // 通过变量调用匿名函数

	// 自执行函数， 匿名函数定义完加()直接执行
	// func(x, y int) {
	// 	fmt.Println(x + y)
	// }(10, 20)
	// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
	// var f = adder()
	// fmt.Println(f(10)) // 10
	// fmt.Println(f(20)) // 30
	// fmt.Println(f(30)) // 60

	// f1 := adder2(10)
	// fmt.Println(f1(40)) // 40
	// fmt.Println(f1(50)) // 90
	// jpgFunc := makeSuffixFunc(".jpg")
	// txtFunc := makeSuffixFunc(".txt")
	// fmt.Println(jpgFunc("test"))
	// fmt.Println(txtFunc("test"))
	// f1, f2 := bcalc(10)
	// fmt.Println(f1(1), f2(2)) //11 9
	// fmt.Println(f1(3), f2(4)) // 12 8
	// fmt.Println(f1(5), f2(6)) // 13	7
	// defer 语句
	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")
	// a := deferf4()
	// fmt.Println(a)
	// x := 1
	// y := 2
	// defer deferf5("AA", x, deferf5("A", x, y))
	// x = 10
	// defer deferf5("BB", x, deferf5("B", x, y))
	// y = 20

	// funcA()
	// funcB()
	// funcC()

}

// panic/recover
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出现了panic错误， 可以通过recover恢复过来
		if err != nil {
			fmt.Println(err)
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func deferf5(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func deferf4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func deferf3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func deferf2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func deferf1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func bcalc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 高阶函数--函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 高阶函数--函数作为参数
func gadd(x, y int) int {
	return x + y
}

func gcalc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) // 变量i只在当前for语句块中生效
	}
	// fmt.Println(i) // 此处无法使用变量i
}

func testLocalVar2(x, y int) {
	fmt.Println(x, y) // 函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 // 变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z) // 此处无法使用变量z
}

func testNum() {
	//如果局部变量和全局变量重名，优先访问局部变量。
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量

}

func testLocalVar() {
	// 定义一个函数局部变量x, 仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) // 函数中可以访问全局变量num
}

// 定义一个带参数，带返回值的函数
func intSum(x int, y int) int {
	return x + y
}

// 定义一个不带参数，不带返回值的函数
func sayHello() {
	fmt.Println("Hello 沙河")
}

// 定义一个可变参数 ...
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
