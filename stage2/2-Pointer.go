package main

import "fmt"

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

func main2() {
	initVal := []int{1, 2, 3}
	fmt.Println("value 1: ", initVal)
	multipleBy2(&initVal)
	fmt.Println("value 2: ", initVal)
}

func multipleBy2(param *[]int) {

	for i := range *param {
		(*param)[i] *= 2
	}

	// Q： 指针需要整体转换为 数组之后，才可以具体操作---加括号
	for i := 0; i < len(*param); i++ {
		// fmt.Println("value 2: ", param)
		// *param[i] = param[i] * 2
		// param[i] = param[i] * 2

		(*param)[i] *= 2
	}
}

// Q: 如何判断接受的是指针，还是数组呢？
// A：方括号在最前 → 数组；星号在最前 → 指针；方括号里没数字 → 切片。

// Q： 什么类型的变量，在引用方法内修改之后，不影响在引用方法外的值 ?
// A:  值接收者永远复制；指针接收者要改本体，但前提是你拿得到“本体”的地址。如果编译器只能先复制再取地址，那最终还是“改了副本”，外面就不受影响。
