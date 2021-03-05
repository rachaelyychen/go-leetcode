package main

type NumArray struct {
	Nums []int
}

func NumArrayConstructor(nums []int) NumArray {
	return NumArray{Nums: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	var res int
	for k := i; k <= j; k++ {
		res += this.Nums[k]
	}
	return res
}
