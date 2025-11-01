package main

import (
	"sort"
)

func main6() {

}

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

// 示例 1：
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

// 示例 2：
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 示例 3：
// 输入：intervals = [[4,7],[1,4]]
// 输出：[[1,7]]
// 解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。

func merge(intervals [][]int) [][]int {
	// slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}
	res = append(res, []int{left, right}) // 将最后一个区间放入
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//
// Q1：如果使用遍历的话，可能产生无限次遍历的问题； ---A：先排序
// Q2: 如何在遍历的时候，进行合并？                ---A: 遍历的同时，需要另外一个指针； 遍历的过程中添加程序，移动指针 & 合并并赋值添加到新的集合数据中
