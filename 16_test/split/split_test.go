package split

import (
	"reflect"
	"testing"
)

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// func benchmarkFib(b *testing.B, n int) {
// 	for i := 0; i < b.N; i++ {
// 		Fib(n)
// 	}
// }

// func benchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
// func benchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
// func benchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
// func benchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
// func benchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
// func benchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

// // 基准测试示例
// func BenchmarkSplik(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Split("沙河有沙又有河", "沙")
// 	}
// }

// func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T的参数
// 	got := Split("a:b:c", ":")         // 程序输出的结果
// 	want := []string{"a", "b", "c"}    // 期望的结果
// 	if !reflect.DeepEqual(want, got) { // 因为slice不能直接比较，借助反射包中的方法比较
// 		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
// 	}
// }

// func TestMoreSplit(t *testing.T) {
// 	got := Split("abcd", "bc")
// 	want := []string{"a", "d"}
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("excepted:%v, got:%v", want, got)
// 	}
// }

// // 测试组Slice
// func TestSplit(t *testing.T) {
// 	// 定义一个测试用例类型
// 	type test struct {
// 		input string
// 		sep   string
// 		want  []string
// 	}
// 	// 定义一个存储测试用例的切片
// 	tests := []test{
// 		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
// 		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
// 		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
// 		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
// 	}
// 	// 遍历切片，逐一执行测试用例
// 	for _, tc := range tests {
// 		got := Split(tc.input, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
// 		}
// 	}
// }

// // 测试组，map实现子测试
// func TestSplit(t *testing.T) {
// 	type test struct { // 定义test结构体
// 		input string
// 		sep   string
// 		want  []string
// 	}
// 	tests := map[string]test{ // 测试用例使用map存储
// 		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
// 		"wrong_sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
// 		"more_sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
// 		"leading_sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
// 	}
// 	for name, tc := range tests {
// 		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
// 			got := Split(tc.input, tc.sep)
// 			if !reflect.DeepEqual(got, tc.want) {
// 				t.Errorf("excepted:%#v, got:%#v", tc.want, got) // 将测试用例的name格式化输出
// 			}
// 		})

// 	}
// }
