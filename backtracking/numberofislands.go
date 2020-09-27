package main

import "fmt"

func main() {
	//[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
	grid := make([][]byte, 4)
	grid[0] = []byte{'1', '1', '1', '1', '0'}
	grid[1] = []byte{'1', '1', '0', '1', '0'}
	grid[2] = []byte{'1', '1', '0', '0', '0'}
	grid[3] = []byte{'0', '0', '0', '0', '0'}
	fmt.Println(numIslands(grid))
	fmt.Println(numIslands_UpdateSameGrid(grid))

	grid2 := make([][]byte, 4)
	grid2[0] = []byte{'1', '1', '0', '0', '0'}
	grid2[1] = []byte{'1', '1', '0', '0', '0'}
	grid2[2] = []byte{'0', '0', '1', '0', '0'}
	grid2[3] = []byte{'0', '0', '0', '1', '1'}
	fmt.Println(numIslands(grid2))
	fmt.Println(numIslands_UpdateSameGrid(grid2))
}

/*-------------------------------------------------------------------
DFS soln with no extra grid where we update the same grid with 1 => 0 once we visit 1
*/
var Rows, Cols int

func numIslands_UpdateSameGrid(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	Rows, Cols = len(grid), len(grid[0])
	cnt := 0
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			if grid[i][j] == '0' {
				continue
			}

			cnt++
			dfs(grid, i, j)
		}
	}
	return cnt
}

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || i > Rows-1 || j < 0 || j > Cols-1 {
		return
	}

	// fmt.Println("i, j  = ", i, j)
	if grid[i][j] == '0' {
		return
	}

	//UPDATE the grid value to '0'
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
}

/*-------------------------------------------------------------------
Use a visited grid - uses extra memory
idea is to make a visited grid marking each island with value `true` as we find an island & increasing the count
*/
func numIslands(grid [][]byte) int {
	v := make([][]bool, len(grid))
	// CANNOT init 2d array like below
	// for _, val := range v {
	//     fmt.Println(val)
	//     val = make([]bool, len(grid[0]))
	//     fmt.Println(val)
	// }
	// fmt.Println(v)

	for i, _ := range v {
		v[i] = make([]bool, len(grid[0]))
	}

	islandCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			// if visited move ahead
			if v[i][j] {
				continue
			}

			// if its an island, i-e grid[i][j] == '1'
			if grid[i][j] == '1' {
				islandCount++
				buildIsland(grid, i, j, v)
			}
		}
	}
	return islandCount
}

var one rune = '1'

/*
DFS to mark the visited grid = true for all 1s
*/
func buildIsland(grid [][]byte, i int, j int, v [][]bool) {
	// if i, j outside the grid return
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 {
		return
	}

	// if already visited (since we check all directions each time) OR grid[i][j] == '0' i-e we find water side return
	if v[i][j] || grid[i][j] == '0' {
		return
	}

	//set v[i][j] = true
	v[i][j] = true

	//check for land in all sides of the island
	//right
	buildIsland(grid, i+1, j, v)

	//down
	buildIsland(grid, i, j-1, v)

	//left
	buildIsland(grid, i-1, j, v)

	//top
	buildIsland(grid, i, j+1, v)
}
