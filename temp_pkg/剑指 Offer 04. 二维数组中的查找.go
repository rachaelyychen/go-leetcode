package temp_pkg

import "fmt"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/18/21 11:44 AM
**/

// 二维数组的每行、每列都是递增的，查找某个数字是否存在于数组中。
// 时间复杂度O(n)，空间复杂度O(1)的解法：从右上角的数字开始查找，等于被查数则返回；
// 大于被查数则剔除所在列，小于被查数则剔除所在行，不断缩小查找范围。

func main() {
	fmt.Println(Search2DArray([][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}, 7))
}

func Search2DArray(arr [][]int, num int) bool {
	row, col := len(arr), len(arr[0])
	i, j := 0, col-1
	for {
		if i >= row || j < 0 {
			return false
		}
		if arr[i][j] == num {
			return true
		}
		if arr[i][j] > num {
			j -= 1
		} else {
			i += 1
		}
	}
}
