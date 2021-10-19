package temp_pkg

import (
	"fmt"
	"strings"
)

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/18/21 2:29 PM
**/

func main() {
	fmt.Println(replaceSpaceInStr("Here we are."))
	a1, a2 := [10]int{0, 2, 7, 8, 9}, [5]int{1, 3, 4, 5, 6}
	merge2Array(&a1, &a2)
	fmt.Println(a1)
}

// 实现一个函数，将字符串中的空格都替换为"%20"。
// strings包下的Replace()、ReplaceAll()方法，也是通过拼接字符串实现替换的。
func replaceSpaceInStr(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// 合并两个排序数组：有两个增序数组a1和a2，a1空间充足，将a2的元素都插入a1，保持a1增序。
// 时间复杂度O(n+m)，空间复杂度O(1)的解法：使用三个下标索引，第一个指向a1最后一个元素的位置，第二个指向a2最后一个元素的位置，第三个指向a1+a2最后一个元素的位置；
// 第一和第二个索引负责获取元素，第三个索引负责从后往前构建新数组。
// 注：数组从前往后插入需要重复移动大量元素，换个角度从后往前插入，可以减少移动次数。
func merge2Array(a1 *[10]int, a2 *[5]int) {
	i, j, k := 4, 4, 9
	for {
		if i < 0 && j < 0 {
			return
		}
		if i < 0 {
			a1[k] = a2[j]
			k -= 1
			j -= 1
			continue
		}
		if j < 0 {
			a1[k] = a1[i]
			k -= 1
			i -= 1
			continue
		}
		if a1[i] > a2[j] {
			a1[k] = a1[i]
			k -= 1
			i -= 1
		} else {
			a1[k] = a2[j]
			k -= 1
			j -= 1
		}
	}
}
