package mian

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
	fmt.Println(find(grid, grid[x][y], x, y, z, w))
}
func find(grid [][]int, cur int, i, j, z, w int) int {
	count := 0
	a := 4
	for a > 0 {
		// 上
		if a == 4 {
			if i != 0 {
				if grid[i-1][j] > cur {
					if i-1 == z && j == w {
						count++
						return count
					}
					count += find(grid, grid[i-1][j], i-1, j, z, w)
				}
			}
		}
		// 下
		if a == 3 {
			if i != len(grid)-1 {
				if grid[i+1][j] > cur {
					if i+1 == z && j == w {
						count++
						return count
					}
					count += find(grid, grid[i+1][j], i+1, j, z, w)
				}
			}
		}
		// 左右
		if a == 2 {
			if j != 0 {
				if grid[i][j-1] > cur {
					if i == z && j+1 == w {
						count++
						return count
					}
					count += find(grid, grid[i][j-1], i, j-1, z, w)
				}
			}
		}
		if a == 1 {
			if j != len(grid[0])-1 {
				if grid[i][j+1] > cur {
					if i == z && j+1 == w {
						count++
						return count
					}
					count += find(grid, grid[i][j+1], i, j+1, z, w)
				}
			}
		}
		a--
	}
	return count
}
