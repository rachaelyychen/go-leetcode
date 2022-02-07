package main

// 时间复杂度O(n)的解法:单调减的双端队列
// 遍历数组,考虑将每个元素加入队列:首先队列的队首元素是当前窗口的最大值,队列内元素单调减;
// 如果当前元素大于队首元素,就把整个队列清空,再加入当前元素;
// 如果当前元素小于等于队首元素,就把整个队列出队,再把大于等于当前元素的重新入队,最后把当前元素放在队尾,这里首先要判断下队首元素是否已经从窗口移除

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var arr []int
	for i := 0; i < len(nums); i++ {
		if len(arr) == 0 {
			arr = append(arr, nums[i])
		} else {
			if nums[i] > arr[0] {
				arr = arr[0:0]
				arr = append(arr, nums[i])
			} else {
				if i >= k && arr[0] == nums[i-k] {
					if k == 1 {
						arr = arr[:0]
					} else {
						arr = arr[1:]
					}
				}
				var index int
				for j := 0; j+1 <= len(arr); j++ {
					if arr[j] >= nums[i] && (j+1 == len(arr) || arr[j+1] < nums[i]) {
						index = j + 1
					}
				}
				if len(arr) > 0 {
					arr = arr[:index]
				}
				arr = append(arr, nums[i])
			}
		}
		if i >= k-1 {
			res = append(res, arr[0])
		}
	}
	return res
}
