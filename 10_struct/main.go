package main

import "fmt"

// // 类型定义
// type NewInt int

// // 类型别名
// type MyInt = int

// type Person struct {
// 	name string
// 	city string
// 	age  int8
// }

// // 结构体匿名字段
// type Person2 struct {
// 	string
// 	int
// }

// type test struct {
// 	a int8
// 	b int8
// 	c int8
// 	d int8
// }

// type student struct {
// 	name string
// 	age  int
// }

// // 嵌套结构体
// type Address struct {
// 	Province string
// 	City     string
// }

// // User 用户结构体
// type User struct {
// 	Name    string
// 	Gender  string
// 	Address Address
// }

// // 嵌套匿名字段
// type Adderss2 struct {
// 	Province string
// 	City     string
// }

// // User 用户结构体
// type User2 struct {
// 	Name     string
// 	Gender   string
// 	Adderss2 // 匿名字段
// }

// // 嵌套结构体的字段名冲突
// // Address 地址结构体
// type Address3 struct {
// 	Province   string
// 	City       string
// 	CreateTime string
// }

// // Email 邮箱结构体
// type Email3 struct {
// 	Account    string
// 	CreateTime string
// }

// //User 用户结构体
// type User3 struct {
// 	Name   string
// 	Gender string
// 	Address3
// 	Email3
// }

// // 结构体的继承
// // Animal 动物
// type Animal struct {
// 	name string
// }

// func (a *Animal) move() {
// 	fmt.Printf("%s会动! \n", a.name)
// }

// //Dog 狗
// type Dog struct {
// 	Feet    int8
// 	*Animal // 通过嵌套匿名结构体实现继承
// }

// func (d *Dog) wang() {
// 	fmt.Printf("%s会汪汪汪~\n", d.name)
// }

// // 结构体与JSON序列化
// // Student 学生
// type Student struct {
// 	ID     int
// 	Gender string
// 	Name   string
// }

// // Class 班级
// type Class struct {
// 	Title    string
// 	Students []*Student
// }

// // 结构体标签（Tag）
// // Student 学生
// type Student struct {
// 	ID     int    `json:"id"` // 通过指定tag实现json序列化该字段时的key
// 	Gender string //json序列化是默认使用字段名作为key
// 	name   string // 私有不能被json包访问
// }

// 结构体和方法补充知识点
type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	// p.dreams = dreams
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	p1 := Person{name: "通州毛总", age: 10}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)
	// s1 := Student{
	// 	ID:     1,
	// 	Gender: "男",
	// 	name:   "通州毛总",
	// }
	// data, err := json.Marshal(s1)
	// if err != nil {
	// 	fmt.Println("json marshal failed!")
	// 	return
	// }
	// fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
	// c := &Class{
	// 	Title:    "101",
	// 	Students: make([]*Student, 0, 200),
	// }

	// for i := 0; i < 10; i++ {
	// 	stu := &Student{
	// 		Name:   fmt.Sprintf("stu%02d", i),
	// 		Gender: "男",
	// 		ID:     i,
	// 	}
	// 	c.Students = append(c.Students, stu)
	// }
	// fmt.Println(c)
	// fmt.Printf("%T\n", c)

	// // JSON序列化，结构体--->JSON格式的字符串
	// data, err := json.Marshal(c)
	// if err != nil {
	// 	fmt.Println("json marsha1 failed")
	// 	return
	// }
	// fmt.Printf("json:%s\n", data)
	// // JSON反序列化：json格式的字符串--->结构体
	// str := `{"Title":"101", "Students":[{"ID":0, "Gender":"男", "Name":"stu00"},{"ID":1,"Gender":"男", "Name":"stu01"},{"ID":2,"Gender":"男", "Name":"stu02"},{"ID":3,"Gender":"男", "Name":"stu03"},{"ID":4,"Gender":"男", "Name":"stu04"},{"ID":5,"Gender":"男", "Name":"stu05"},{"ID":6,"Gender":"男", "Name":"stu06"},{"ID":7,"Gender":"男", "Name":"stu07"},{"ID":8,"Gender":"男", "Name":"stu08"},{"ID":9,"Gender":"男", "Name":"stu09"}]}`
	// c1 := &Class{}
	// fmt.Println(c1)
	// fmt.Printf("%T\n", c1)
	// err = json.Unmarshal([]byte(str), c1)
	// if err != nil {
	// 	fmt.Println("json unmarshal failed!")
	// 	return
	// }
	// fmt.Printf("%#v\n", c1)
	// d1 := &Dog{
	// 	Feet: 4,
	// 	Animal: &Animal{
	// 		// 注意嵌套的是结构体指针
	// 		name: "乐乐",
	// 	},
	// }
	// d1.wang()
	// d1.move()
	// var user3 User3
	// user3.Name = "通州毛总"
	// user3.Gender = "男"
	// // user3.CreateTime = "2019"  // ambiguous selector user3.CreateTime
	// user3.Address3.CreateTime = "2000" // 指定Address3结构体中的CreateTime
	// user3.Email3.CreateTime = "2001"   // 指定Email结构体中的CreteTime
	// fmt.Printf("%#v \n", user3)        // main.User3{Name:"通州毛总", Gender:"男", Address3:main.Address3{Province:"", City:"", CreateTime:"2000"}, Email3:main.Email3{Account:"", CreateTime:"2001"}}
	// var user2 User2
	// user2.Name = "通州毛总"
	// user2.Gender = "男"
	// user2.Adderss2.Province = "安徽"   // 匿名字段默认使用类型名作为字段名
	// user2.City = "宿州"                // 匿名字段可以省略
	// fmt.Printf("user2=%#v\n", user2) // user2=main.User2{Name:"通州毛总", Gender:"男", Adderss2:main.Adderss2{Province:"安徽", City:"宿州"}}
	// user1 := User{
	// 	Name:   "通州毛总",
	// 	Gender: "男",
	// 	Address: Address{
	// 		Province: "安徽",
	// 		City:     "宿州",
	// 	},
	// }
	// fmt.Printf("user1=%#v\n", user1) // user1=main.User{Name:"通州毛总", Gender:"男", Address:main.Address{Province:"安徽", City:"宿州"}}

	// var a NewInt
	// var b MyInt

	// fmt.Printf("type of a:%T\n", a) // type of a:main.NewInt
	// fmt.Printf("type of b:%T\n", b) // type of b:int
	// var p1 person
	// p1.name = "沙河娜扎"
	// p1.city = "北京"
	// p1.age = 18
	// fmt.Printf("p1=%v\n", p1)
	// fmt.Printf("p1=%#v\n", p1)

	// // 匿名结构体
	// var user struct {
	// 	Name string
	// 	Age  int
	// }
	// user.Name = "小王子"
	// user.Age = 18
	// fmt.Printf("%#v\n", user)

	// // 创建指针类型结构体
	// var p2 = new(person)
	// fmt.Printf("%T\n", p2)     // *main.person
	// fmt.Printf("p2=%#v\n", p2) // p2=&main.person{name:"", city:"", age:0}

	// //取结构体的地址实例化
	// p3 := &person{}
	// fmt.Printf("%T\n", p3)     // *main.person
	// fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"", city:"", age:0}
	// p3.name = "七米"
	// p3.age = 30
	// p3.city = "成都"
	// fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"七米", city:"成都", age:30}

	// // 结构体初始化
	// var p4 person
	// fmt.Printf("p4=%#v\n", p4) // p4=main.person{name:"", city:"", age:0}

	// // 使用键值对初始化
	// p5 := person{
	// 	name: "小王子",
	// 	city: "北京",
	// 	age:  18,
	// }
	// fmt.Printf("p5=%#v\n", p5) // p5=main.person{name:"小王子", city:"北京", age:18}

	// // 也可以对结构体指针进行键值对初始化
	// p6 := &person{
	// 	name: "大王子",
	// 	city: "北京",
	// 	age:  19,
	// }
	// fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"大王子", city:"北京", age:19}

	// p7 := &person{
	// 	city: "北京",
	// }
	// fmt.Printf("p7=%#v\n", p7) // p7=&main.person{name:"", city:"北京", age:0}

	// // 1.必须初始化结构体的所有字段。
	// // 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// // 3.该方式不能和键值初始化方式混用。
	// p8 := &person{
	// 	"通州毛总",
	// 	"北京",
	// 	28,
	// }
	// fmt.Printf("p8=%#v\n", p8) // p8=&main.person{name:"通州毛总", city:"北京", age:28}

	// // 结构体内存布局
	// n := test{
	// 	1, 2, 3, 4,
	// }
	// fmt.Printf("n.a %p\n", &n.a) // n.a 0xc0000a0068
	// fmt.Printf("n.b %p\n", &n.b) // n.b 0xc0000a0069
	// fmt.Printf("n.c %p\n", &n.c) // n.c 0xc0000a006a
	// fmt.Printf("n.d %p\n", &n.d) // n.d 0xc0000a006b

	// var v struct{}
	// fmt.Println(unsafe.Sizeof(v)) // 0

	// m := make(map[string]*student)
	// stus := []student{
	// 	{name: "通州", age: 18},
	// 	{name: "毛毛", age: 23},
	// 	{name: "毛总", age: 9000},
	// }

	// for _, stu := range stus {
	// 	m[stu.name] = &stu
	// }

	// for k, v := range m {
	// 	fmt.Println(k, "=>", v.name)
	// }

	// // 调用构造函数
	// p9 := newPerson("张三", "北京", 90)
	// fmt.Printf("%#v\n", p9) // &main.person{name:"张三", city:"北京", age:90}
	// p9.Dream()
	// p9.SetAge(10)           // 修改结构体年龄
	// fmt.Printf("%#v\n", p9) // &main.Person{name:"张三", city:"北京", age:10}
	// p9.SetAge2(80)
	// fmt.Printf("%#v\n", p9) // &main.Person{name:"张三", city:"北京", age:10}

	// // 什么时候应该使用指针类型接收者
	// // 需要修改接收者中的值
	// // 接收者是拷贝代价比较大的大对象
	// // 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

	// var m1 NewInt
	// m1.SayHello() // Hello, 我是一个int.
	// m1 = 100
	// fmt.Printf("%#v %T\n", m1, m1) // 100 main.NewInt

	// p1 := Person2{
	// 	"毛总",
	// 	18,
	// }
	// fmt.Printf("%#v\n", p1)
	// fmt.Println(p1.string, p1.int)
	// 注意：这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
}

// // 构造函数
// func newPerson(name, city string, age int8) *Person {
// 	return &Person{
// 		name: name,
// 		city: city,
// 		age:  age,
// 	}
// }

// // Dream Person做梦的方法
// func (p Person) Dream() {
// 	fmt.Printf("%s的梦想是学好GO语言！\n", p.name)
// }

// // SetAge 设置P的年龄， 使用指针接收者
// func (p *Person) SetAge(newAge int8) {
// 	p.age = newAge
// }

// // SetAge2 设置p的年龄，使用值接收者
// func (p Person) SetAge2(newAge int8) {
// 	p.age = newAge
// }

// // 任意类型添加方法
// // SayHello 为MyInt添加一个SayHello的方法
// func (m NewInt) SayHello() {
// 	fmt.Println("Hello, 我是一个int.")
// }
