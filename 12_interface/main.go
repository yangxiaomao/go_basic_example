package main

import "fmt"

// // 定义一个dog 结构体
// type dog struct{}

// // 定义一个cat 结构体
// type cat struct{}

// // 定义一个Mover 接口
// type Mover interface {
// 	move()
// }

// // 定义一个Mover 接口
// type Mover2 interface {
// 	move2()
// }

// // 值接收者实现接口
// func (d dog) move() {
// 	fmt.Println("狗会动")
// }

// // 指针接收者实现接口
// func (d *dog) move2() {
// 	fmt.Println("狗会动")
// }

// // 定义一个People 接口
// type People interface {
// 	Speak(string) string
// }

// type Student struct{}

// func (s *Student) Speak(t string) (talk string) {
// 	if t == "sg" {
// 		talk = "你是一个大帅哥"
// 	} else {
// 		talk = "您好"
// 	}
// 	return
// }
// // Sayer 接口
// type Sayer interface {
// 	say()
// }

// // Mover 接口
// type Mover interface {
// 	move()
// }

// // dog既可以实现Sayer接口， 也可以实现Mover接口
// type dog struct {
// 	name string
// }

// // 实现Sayer接口
// func (d dog) say() {
// 	fmt.Printf("%s会叫汪汪汪\n", d.name)
// }

// // 实现Mover接口
// func (d dog) move() {
// 	fmt.Printf("%s会动\n", d.name)
// }

// 不同类型实现同一接口
// Mover 接口
// type Mover interface {
// 	move()
// }

// // 定义dog 结构体
// type dog struct {
// 	name string
// }

// // 定义car 结构体
// type car struct {
// 	brand string
// }

// // dog类型实现Mover接口
// func (d dog) move() {
// 	fmt.Printf("%s会跑\n", d.name)
// }

// // car类型实现Mover接口
// func (c car) move() {
// 	fmt.Printf("%s速度70迈\n", c.brand)
// }

// // WashingMachine 洗衣机
// type WashingMachine interface {
// 	wash()
// 	dry()
// }

// // 甩干器
// type dryer struct{}

// // 实现WashingMachine接口的dry()方法
// func (d dryer) dry() {
// 	fmt.Println("甩一甩")
// }

// // 海尔洗衣机
// type haier struct {
// 	dryer // 嵌入甩干器
// }

// // 实现WashingMachine接口的wash()方法
// func (h haier) wash() {
// 	fmt.Println("洗刷刷")
// }

// // Sayer 接口
// type Sayer interface {
// 	say()
// }

// // Mover 接口
// type Mover interface {
// 	move()
// }

// // 接口嵌套
// type animal interface {
// 	Sayer
// 	Mover
// }

// type cat struct {
// 	name string
// }

// func (c cat) say() {
// 	fmt.Println("喵喵喵")
// }

// func (c cat) move() {
// 	fmt.Println("猫会动")
// }

// // 空接口作为函数参数
// func show(a interface{}) {
// 	fmt.Printf("type:%T value:%v\n", a, a)
// }

// 上面的示例中如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现：
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type! ")
	}
}
func main() {
	justifyType("毛总")
	// // 空接口数据类型断言
	// var x interface{}
	// x = "毛总"
	// v, ok := x.(string)
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("类型断言失败")
	// }
	// // 空接口作为map值
	// var studentInfo = make(map[string]interface{})
	// studentInfo["name"] = "毛总"
	// studentInfo["age"] = 18
	// studentInfo["married"] = false
	// fmt.Printf("%#v", studentInfo)
	// show("Hello 毛总")
	// // 定义一个空接口x
	// var x interface{}
	// s := "Hello 毛总"
	// x = s
	// fmt.Printf("type:%T value:%v\n", x, x)
	// i := 100
	// x = i
	// fmt.Printf("type:%T value:%v\n", x, x)
	// b := true
	// x = b
	// fmt.Printf("type:%T value:%v\n", x, x)

	// var x animal
	// x = cat{name: "花花"}
	// x.move()
	// x.say()
	// var x WashingMachine
	// var a = haier{}
	// x = a
	// x.wash()
	// x.dry()
	// var x Mover
	// var a = dog{name: "旺财"}
	// var b = car{brand: "保时捷"}
	// x = a
	// x.move()
	// x = b
	// x.move()
	// var x Sayer
	// var y Mover
	// var a = dog{name: "旺财"}
	// x = a
	// y = a
	// x.say()
	// y.move()
	// // 定义的Student的方法中接收者是指针形式，所以用 Student{}值类型的用不了
	// var peo People = &Student{}
	// think := "sg"
	// fmt.Println(peo.Speak(think))
	// var x Mover2
	// // var wangcai = dog{} // wangcai是dog类型
	// // x = wangcai         // x不可以接收dog类型
	// var fugui = &dog{} //fugui 是*dog类型
	// x = fugui          // x可以接收*dog类型
	// x.move2()
	// var x Mover
	// var wangcai = dog{} // wangcai是dog类型
	// x = wangcai         // x可以接收dog类型
	// x.move()
	// var fugui = &dog{} // fugui是*dog类型
	// x = fugui          // x可以接收*dog类型
	// x.move()
	// //初始化命令行参数
	// flag.Parse()
	// glog.Info("hello, glog")
	// // 退出时调用，确保日志写入文件中
	// glog.Flush()
	// var x Sayer // 声明一个Sayer类型的变量
	// a := cat{}  // 实例化一个cat
	// b := dog{}  // 实例化一个dog
	// x = a       // 可以把cat实例直接赋值给x
	// x.say()     // 喵喵喵
	// x = b       // 可以把dog实例直接赋值给x
	// x.say()     // 汪汪汪

}

// // Sayer 接口
// type Sayer interface {
// 	say()
// }

// // dog实现了 Sayer接口
// func (d dog) say() {
// 	fmt.Println("汪汪汪")
// }

// // cat实现了 Sayer接口
// func (c cat) say() {
// 	fmt.Println("喵喵喵")
// }
