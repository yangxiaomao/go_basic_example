package main

import "fmt"

func main() {
	// var a = [3]int{1, 2, 3}
	// b := arraySum(a)
	// fmt.Println(b)
	// deSlice()
	// deSlice1()
	deSlice2()
}

//使用make()函数构造切片
func deSlice2() {
	// a := make([]int, 2, 10)
	// fmt.Println(a)      // [0 0]
	// fmt.Println(len(a)) // 2
	// fmt.Println(cap(a)) // 10

	// 切片不能直接比较
	// var s1 []int
	// s2 := []int{}
	// s3 := make([]int, 0)
	// fmt.Println(s1)
	// fmt.Println(s2)
	// fmt.Println(s3)
	// fmt.Println(s3 == s1)

	// // 赋值和拷贝
	// s1 := make([]int, 3) //[0 0 0]
	// s2 := s1             //这里将s1直接赋值给s2, s1和s2是同一个底层数组
	// s2[0] = 100
	// fmt.Println(s1) //[100 0 0]
	// fmt.Println(s2) //[100 0 0]

	// // 切片的遍历
	// s := []int{1, 3, 5}
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, s[i])
	// }

	// for index, value := range s {
	// 	fmt.Println(index, value)
	// }

	// append()方法为切片添加元素
	// var s []int
	// s = append(s, 1)
	// s = append(s, 2, 3, 4, 5)
	// s2 := []int{5, 6, 7}
	// s = append(s, s2...)
	// fmt.Println(s)
	// fmt.Println(s2)
	// append()添加元素和切片扩容
	// var numSlice []int
	// for i := 0; i < 10; i++ {
	// 	numSlice = append(numSlice, i)
	// 	fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	// }
	// var citySlice []string
	// // 追加一个元素
	// citySlice = append(citySlice, "北京")
	// // 追加多个元素
	// citySlice = append(citySlice, "上海", "广州", "深圳")
	// // 追加切片
	// a := []string{"重庆", "成都"}
	// citySlice = append(citySlice, a...)
	// fmt.Println(citySlice) // [北京 上海 广州 深圳 重庆 成都]

	// // 使用copy()函数复制切片
	// a := []int{1, 2, 3, 4, 5}
	// b := a
	// fmt.Println(a)
	// fmt.Println(b)
	// b[0] = 1000
	// fmt.Println(a)
	// fmt.Println(b)
	// a := []int{1, 2, 3, 4, 5}
	// c := make([]int, 5, 5)
	// copy(c, a)
	// fmt.Println(a) //[1 2 3 4 5]
	// fmt.Println(c) //[1 2 3 4 5]
	// c[0] = 1000
	// fmt.Println(a) //[1 2 3 4 5]
	// fmt.Println(c) //[1000 2 3 4 5]

	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// // 要删除索引为2的元素
	// a = append(a[:2], a[3:]...)
	// fmt.Println(a)

	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)      // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(len(a)) // 15

}

//简单的切片表达式
func slice1() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}

//简单切片表达式
func deSlice1() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[1:3]
	// fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	// fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	// a[2:]  // 等同于 a[2:len(a)]
	// a[:3]  // 等同于 a[0:3]
	// a[:]   // 等同于 a[0:len(a)]
}

// 切片的定义
func deSlice() {
	// var name []T
	// 声明切片类型
	var a []string              //声明一个字符类型的切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
}

// 因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。 例如：
func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
