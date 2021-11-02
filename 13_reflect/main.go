package main

// func reflectType(x interface{}) {
// 	v := reflect.TypeOf(x)
// 	fmt.Printf("type:%v\n", v)
// }
// type myInt int64

// func reflectType(x interface{}) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
// }

// // 通过反射获取值
// func reflectValue(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Int64:
// 		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
// 		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
// 	case reflect.Float32:
// 		// v.Float()从反射中获取浮点型的原始值， 然后通过float32()强制类型转换
// 		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
// 	case reflect.Float64:
// 		// v.Float()从反射中获取浮点型的原始值， 然后通过float64()强制类型转换
// 		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
// 	}
// }

// func reflectSetValue1(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	if v.Kind() == reflect.Int64 {
// 		v.SetInt(200) // 修改的是副本， reflect包会引发panic
// 	}
// }

// func reflectSetValue2(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	// 反射中使用 Elem()方法获取指针对应的值
// 	if v.Elem().Kind() == reflect.Int64 {
// 		v.Elem().SetInt(200)
// 	}
// }

// type student struct {
// 	Name  string `json:"name"`
// 	Score int    `json:score`
// }

func main() {
	// stu1 := student{
	// 	Name:  "毛总",
	// 	Score: 90,
	// }

	// t := reflect.TypeOf(stu1)
	// fmt.Println(t.Name(), t.Kind()) // student struct
	// // 通过for循环遍历结构体的所有字段信息
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))

	// }
	// // 通过字段名获取指定结构体字段信息
	// if scoreField, ok := t.FieldByName("Score"); ok {
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	// }
	// // IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	// // *int类型空指针
	// var a *int
	// fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// // nil值
	// fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid)
	// // 实例化一个匿名结构体
	// b := struct{}{}
	// // 尝试从结构体中查找“abc”字段
	// fmt.Println("不存在的结构体成员：", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// // 尝试从结构体中查找“abc”方法
	// fmt.Println("不存在的结构体方法：", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// // map
	// c := map[string]int{}
	// // 尝试从map中查找一个不存在的键
	// fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("毛总")).IsValid())

	// var a int64 = 100
	// // reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	// reflectSetValue2(&a)
	// fmt.Println(a)
	// var a float32 = 3.14
	// var b int64 = 100
	// reflectValue(a) // type is float32, value is 3.140000
	// reflectValue(b) // type is int64, value is 100
	// // 将int类型的原始值转换为reflect.Value类型
	// c := reflect.ValueOf(10)
	// fmt.Printf("type c :%T\n", c) // type c :reflect.Value
	// var a *float32 // 指针
	// var b myInt    // 自定义类型
	// var c rune     // 类型别名
	// reflectType(a) // type: kind:ptr
	// reflectType(b) // type:myInt kind:int64
	// reflectType(c) // type:int32 kind:int32

	// type person struct {
	// 	name string
	// 	age  int
	// }
	// type book struct{ title string }
	// var d = person{
	// 	name: "毛总",
	// 	age:  18,
	// }
	// var e = book{title: "<跟毛总学GO语言>"}
	// reflectType(d) // type:person kind:struct
	// reflectType(e) // type:book kind:struct
	// var a float32 = 3.14
	// reflectType(a) // type:float32
	// var b int64 = 100
	// reflectType(b) // type:int64
}
