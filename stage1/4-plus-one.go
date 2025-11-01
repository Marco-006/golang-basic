package main

// 难度：简单
// 考察：数组操作、进位处理
// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

func main4(){
	
	
}



func plusOne(digits []int) []int {
    n := len(digits)
    for i := n - 1; i >= 0; i-- {
        if digits[i] != 9 {
            digits[i]++
            for j := i + 1; j < n; j++ {
                digits[j] = 0
            }
            return digits
        }
    }
    // digits 中所有的元素均为 9

    digits = make([]int, n+1)
    digits[0] = 1
    return digits
}


func test1(){
	// nums := [...]uint64{1,2,3,4,5,6}

	// //先转换为  string
	// var str string 
	// for _, v := range nums {
	// 	str = append(str, string(v))
	// }

	// //sring转换为 uint64 ；  +1
	// //拆分为数组
}

