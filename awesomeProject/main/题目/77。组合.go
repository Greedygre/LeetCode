package 题目

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//
//输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	rList := make([]int, 0, 100)
	res := deepCombine(n, k, 1, [][]int{}, rList)

	return res
}

func deepCombine(n, k, start int, res [][]int, list []int) [][]int {
	if len(list) == k {
		v := make([]int, 0, len(list))
		//copy(v, list)
		v = append(v, list...)
		res = append(res, v)
		return res
	}
	maxN := (n - (k - len(list))) + 1
	for i := start; i <= maxN; i++ {
		list = append(list, i)
		res = deepCombine(n, k, i+1, res, list)
		list = list[:len(list)-1]
	}

	return res
}
