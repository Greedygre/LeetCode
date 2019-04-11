package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	m := 0
	fmt.Scan(&m)
	grid := make([][]int, 0)
	for i := 0; i < n; i++ {
		cur := []int{}
		for j := 0; j < m; j++ {
			c := 0
			fmt.Scan(&c)
			cur = append(cur, c)
		}
		grid = append(grid, cur)
	}
	x := 0
	fmt.Scan(&x)
	y := 0
	fmt.Scan(&y)
	z := 0
	fmt.Scan(&z)
	w := 0
	fmt.Scan(&w)
	fmt.Println(find(grid, grid[x][y], x, y, z, w, -1))
}
func find(grid [][]int, cur int, i, j, z, w int, zhuangtai int) int {
	if i == z && j == w {
		return 1
	}
	count := 0
	a := 4
	for a > 0 {
		// 上
		if a == 4 && zhuangtai != 2 {
			if i != 0 {
				if grid[i-1][j] > cur {
					count += find(grid, grid[i-1][j], i-1, j, z, w, 1)
				}
			}
		}
		// 下
		if a == 3 && zhuangtai != 1 {
			if i != len(grid)-1 {
				if grid[i+1][j] > cur {
					count += find(grid, grid[i+1][j], i+1, j, z, w, 2)
				}
			}
		}
		// 左右
		if a == 2 && zhuangtai != 4 {
			if j != 0 {
				if grid[i][j-1] > cur {
					count += find(grid, grid[i][j-1], i, j-1, z, w, 3)
				}
			}
		}
		if a == 1 && zhuangtai != 3 {
			if j != len(grid[0])-1 {
				if grid[i][j+1] > cur {
					count += find(grid, grid[i][j+1], i, j+1, z, w, 4)
				}
			}
		}
		a--
	}
	return count
}

//4 5
//0 1 0 0 0
//0 2 3 0 0
//0 3 4 5 0
//0 0 7 6 0
//0 1 3 2
