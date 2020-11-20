package main

import "fmt"

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println(s1)
	s2 := s1[3:5]
	fmt.Println(s2)
	s3 := append(s2, 10)
	fmt.Println(s3)
	s4 := append(s3, 11)
	fmt.Println(s4)
	s5 := append(s4, 12)
	fmt.Println(s5, arr)
	return

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1

		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
