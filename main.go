package main

func main() {
	// pic.Show(Pic)
	// fmt.Println(os.Args)
	// fmt.Printf("%T\n", os.Args)
	// fmt.Println(len(os.Args))
	// if len(os.Args) > 1 {
	// 	fmt.Println("Hello World", os.Args[1])
	// }
	// sls := []string{"1", "2", "3", "4", "100", "98"}
	// s5 := make([]string, 0)
	// s5 = append(s5, "100", "101", "103")
	// s6 := s5
	// s6[0] = "1001"
	// fmt.Println(s5)
	// s1 := 1
	// fmt.Println(&s1)
	// s2 := "hello"
	// fmt.Println(&s2)
	// fmt.Println(&sls)

	// a := &sls
	// sls = reverse1(a)
	// fmt.Println(sls)
	// s := 1
	// fmt.Println(s)
	// fmt.Println(Sqrt(2))
	// fmt.Println(math.Sqrt(2))
}

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for math.Abs(x-z*z) > 0.000000001 {
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z
// }

// func string1(s1, s2 string) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}
// 	var num int
// 	for i := len(s1); i >= 0; i-- {
// 		for j := len(s2); j >= 0; j-- {
// 			if s1[i] == s2[j] {
// 				num++
// 			}
// 		}
// 	}

// 	if num == (len(s1) - 1) {
// 		return true
// 	}
// 	return false
// }

// func string2(s1, s2 string) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}

// 	for i := len(s1); i >= 0; i-- {
// 		if strings.LastIndex(s2, string(s1[i])) == -1 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func reverse(s []string) []string {
// 	var n []string
// 	for i := len(s) - 1; i >= 0; i-- {
// 		n = append(n, string(s[i]))
// 	}
// 	return n
// }

// func reverse1(s *[]string) []string {
// 	var n []string
// 	for i := len(*s) - 1; i >= 0; i-- {
// 		n = append(n, (*s)[i])
// 	}
// 	return n
// }

// func Pic(dx, dy int) [][]uint8 {
// 	ans := make([][]uint8, dy)
// 	for y := 0; y < dy; y++ {
// 		row := make([]uint8, dx)
// 		for x := 0; x < dx; x++ {
// 			fx := float64(x)
// 			fy := float64(y)
// 			row[x] = uint8(math.Pow(fx, fy))
// 		}
// 		ans[y] = row
// 	}
// 	return ans
// }
