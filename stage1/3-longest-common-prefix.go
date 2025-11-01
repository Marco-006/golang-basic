package main

// 考察：字符串处理、循环嵌套
// 题目：查找字符串数组中的最长公共前缀

// 示例 1：

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：

// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

func main3() {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	strs3 := []string{"ab", "a"}
	strs4 := []string{"reflower", "flow", "flight"}

	println(longestCommonPrefix(strs1))
	println(longestCommonPrefix(strs2))
	println(longestCommonPrefix(strs3))
	println(longestCommonPrefix(strs4))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 以第 0 个串为基准
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			// 越界或字符不等就结束
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0] // 第 0 个串就是公共前缀

}



// func longestCommonPrefix(strs []string) string {
// 	var result string
// 	for k, v := range strs {
// 		if k == 0 {
// 			result = v
// 		} else {
// 			// 两个字符串，取公共字符串
// 			comonLength := 0
// 			if len(v) < len(strs[k-1]) {
// 				comonLength = len(v)
// 			} else {
// 				comonLength = len(strs[k-1])
// 			}

// 			for i := 0; i < comonLength; i++ {
// 				if v[:i+1] == strs[k-1][:i+1] {
// 					result = string(v[:i+1])
// 				} else {
// 					// result = ""
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

// func commonPrefix(s string) string {
// 	set := make(map[string]int)

// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; i < len(s); i++ {
// 			currentStr := s[i:j]

// 			if _, ok := set[currentStr]; ok {
// 				continue
// 			} else {
// 				set[currentStr] = 1
// 			}
// 		}
// 	}

// 	maxLength := 0
// 	for k, _ := range set {
// 		if len(k) > maxLength {
// 			maxLength = len(k)
// 		}
// 	}

// }
