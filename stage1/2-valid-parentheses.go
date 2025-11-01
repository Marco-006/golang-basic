// 考察：字符串处理、栈的使用
// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

package main

func main2() {
	inputParam := "(){}[]]"
	println(isValidate(inputParam))
}

func isValidate(param string) bool {
	// 不引入额外的数据结构，就使用现有的就可以了；

	// 配对 map
	maps := map[byte]byte{')': '(', '}': '{', ']': '['}

	// 结果容器，最终为空就是有效的；不为空就是无效的
	// restultArray := []byte{}
	// println(restultArray)
	// 切片（动态长度，最常用）
	stack := make([]byte, 0)

	if param == "" {
		return true
	}

	for i := 0; i < len(param); i++ {
		v := param[i]
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if len(stack) > 0 && stack[len(stack)-1] == maps[v] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0

	//1.  用 for range 遍历字符串时，默认按“Unicode 码点”（rune） 逐个迭代，而不是按字节
	//2.  想按字节遍历字符串，用下标切片
	// 3. A. 数组（固定长度）B. 切片（动态长度，最常用）

}
