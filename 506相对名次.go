package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	type ScoreRankIndex struct {
		score int
		rank  int
		index int
	}
	var arr []ScoreRankIndex
	for i := range score {
		arr = append(arr, ScoreRankIndex{score: score[i], index: i})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score > arr[j].score
	})
	for i := range arr {
		arr[i].rank = i+1
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].index < arr[j].index
	})
	var res []string
	for i := range arr {
		if arr[i].rank == 1 {
			res = append(res, "Gold Medal")
		} else if arr[i].rank == 2 {
			res = append(res, "Silver Medal")
		} else if arr[i].rank == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(arr[i].rank))
		}
	}
	return res
}
