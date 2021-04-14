package main

type OrderedStream struct {
	Len, Ptr int
	Bucket   []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Len:    n,
		Ptr:    0,
		Bucket: make([]string, n),
	}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	var result = make([]string, 0)
	this.Bucket[id-1] = value
	if this.Bucket[this.Ptr] != "" {
		i := this.Ptr
		for ; i < this.Len; i++ {
			if this.Bucket[i] == "" {
				break
			}
		}
		result = this.Bucket[this.Ptr:i]
		this.Ptr = i
	}
	return result
}
