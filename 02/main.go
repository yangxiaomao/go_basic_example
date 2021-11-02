package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	t1 := time.Now() // get current time
	var normalNum float64 = 10
	var isZero bool = false
	var price int = 1000000
	all := make([]int, int(normalNum))
	for !isZero {
		all = batchNormal(all, price, normalNum)
		var min_val int
		for _, val := range all {
			if min_val == 0 || val <= min_val {
				min_val = val
			}
		}

		if min_val > 0 {
			isZero = true
		} else {
			isZero = false
		}

	}
	//总金额转换String，并保留两位
	total_price := fmt.Sprintf("%.2f", float64(price/100))
	allString := make([]string, int(normalNum))

	for k, v := range all {
		allString[k] = fmt.Sprintf("%.2f", float64(float64(v)/100))
	}
	fmt.Println(total_price)
	fmt.Println(allString)

	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

	// $total_price = bcdiv($price , 100, 2);
	// foreach($all as $key=>$val){
	//     $all[$key] = bcdiv($val, 100, 2);
	// }

	// dd($total_price,$all);

	// fmt.Println(m)
	// fmt.Println(useLast)

	// fmt.Println(float32(3934534534))

}

func GenerateRangeNum(min, max int) int {
	// fmt.Println(r.Intn(max - min))
	randNum := r.Intn(max-min) + min
	return randNum
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value

}

func normalRand() float64 {
	var useLast bool = false
	var m float64 = 0
	var n float64 = 0
	for !useLast {
		//返回自然对数
		ranges := 10000000
		u := float64(GenerateRangeNum(1, ranges)) / float64(ranges)
		v := float64(GenerateRangeNum(1, ranges)) / float64(ranges)
		s := (math.Sqrt(-2 * math.Log(u)))

		x := s * math.Sin(2*math.Pi*v)
		y := s * math.Cos(2*math.Pi*v)

		// 以下为封装
		//分布的宽度
		width := 1.0
		//分布的锋锐程度， 数字越小，曲线越平。
		step := 2.0
		m = (x/(step*width*2.0) + width/2.0)
		n = (y/(step*width*2.0) + width/2.0)
		//              dump($m,$n);
		if m < width && m > 0 && n < width && n > 0 {
			useLast = true
		}
	}
	return m
}

func batchNormal(all []int, price int, normalNum float64) []int {
	for i := 0; i < price; i++ {
		normalRands := normalRand()
		n := int(normalRands * normalNum)
		all[n]++
		// time.Sleep(1 * time.Second)
	}
	return all
}
