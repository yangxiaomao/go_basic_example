package main

import "fmt"

func main() {
	// a := 10
	// b := &a
	// fmt.Printf("a:%d ptr:%p\n", a, &a) //a:10 ptr:0xc000012090
	// fmt.Printf("b:%p type:%T\n", b, b) //b:0xc000012090 type:*int
	// fmt.Println(&b)                    //0xc000006028

	// //指针取值
	// a := 10
	// b := &a                          // 取变量a的地址， 将指针保存到b中
	// fmt.Printf("type of b:%T\n", b)  // type of b:*int
	// c := *b                          // 指针取值（根据指针去内存取值）
	// fmt.Printf("type of c:%T\n", c)  // type of c:int
	// fmt.Printf("value of c:%v\n", c) // value of c:10

	// a := 10
	// modify1(a)
	// fmt.Println(a) // 10 函数modify1传入x值，没有返回值x域变了
	// modify2(&a)
	// fmt.Println(a) // 100 函数modify2传入x的内存指针，虽然域变了但指向同一内存地址

	// new 和 make
	// var a *int
	// *a = 100
	// fmt.Println(*a)

	// var b map[string]int
	// b["沙河娜扎"] = 100    // 这里会报错，无效的内存地址
	// fmt.Println(b)
	// a := new(int)
	// b := new(bool)
	// fmt.Printf("%T\n", a) // *int
	// fmt.Printf("%T\n", b) //*bool
	// fmt.Println(*a)       //0
	// fmt.Println(*b)       // false

	// var a *int
	// a = new(int)
	// *a = 10
	// fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)

}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
