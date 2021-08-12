package main

func maxArea(height []int) int {
	var res int
	// 首先可以确定的一点是，以某条线作为边的面积最大值，就是另一条线比它高且距离它最远时
	// 那怎么在n时间复杂度找到以每条线为边的最大面积呢： 双指针
	// 头尾各一个指针i、j，如果height[i]<height[j]，表示找到了以height[i]为边的最大面积，接着希望找到以height[j]为边的最大面积，于是移动i
	// height[i]>height[j]同理，相等则都移动一下
	i, j := 0, len(height)-1
	for i < j {
		var t int
		if height[i] < height[j] {
			t = height[i] * (j - i)
			i += 1
		} else if height[i] > height[j] {
			t = height[j] * (j - i)
			j -= 1
		} else {
			t = height[i] * (j - i)
			i += 1
			j -= 1
		}
		if t > res {
			res = t
		}
	}
	return res
}
