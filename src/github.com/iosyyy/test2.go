package iosyyy

import (
	"fmt"
)

func TestVer() {
	var a = 4
	ptr := new(int)
	fmt.Println()
	fmt.Printf("%v", &a)
	fmt.Printf("%v", testVer(*ptr))
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
func testVer(a int) (resA *int) {
	if &a == nil {
		println("a为空")
		return
	}
	resA = &a
	return
}

func sort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if (arr)[i] >= (arr)[j] {
				(arr)[j], (arr)[i] = (arr)[i], (arr)[j]
			}
		}
	}

}

func cut(nums ...int) {
	arrs := make([]int, 5)
	copy(arrs, nums)
	fmt.Println(arrs[:])
}

func init() {
	/*maps := map[int]string{1: "hello world", 5: "test"}
	fmt.Println(maps)

	for key, value := range maps {
		fmt.Println("key: ", key, "value: ", value)
	}

	value, ok := maps[0]
	fmt.Println("value: ", value, "ok: ", ok)
	delete(maps, 0)*/

	//cuts := []int{1, 5, 4641, 651, 1321, 32156, 1, 1321}
	//
	//sort(cuts)
	//
	//num := 1001
	//nums := make([]int, 4)
	//for i := range nums {
	//	nums[i] = num % 10
	//	num /= 10
	//}
	//fmt.Println(nums)
	//cut(1, 2, 3, 4, 5, 6)
	/*name := [5]int{1: 1, 3: 6}
	for _, i2 := range name {
		fmt.Print(i2, " ")
	}
	fmt.Println()
	for _, i2 := range name {
		fmt.Print(i2, " ")
	}
	fmt.Println()

	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))*/

	//a, b := 10, 20
	//swap(&a, &b)
	//fmt.Printf("%d %d", a, b)
}
