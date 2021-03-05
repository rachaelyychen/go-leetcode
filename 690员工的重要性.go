package main

import "fmt"

type Employee struct {
	Id int
	Importance int
	Subordinates []int
}

var employeeMap = make(map[int]*Employee, 0)

func getImportance(employees []*Employee, id int) int {
	for i := range employees {
		employeeMap[employees[i].Id] = employees[i]
	}
	return getImportance2(id)
}

func getImportance2(id int) int {
	employee := employeeMap[id]
	res := employee.Importance
	for i := range employee.Subordinates {
		res += getImportance2(employee.Subordinates[i])
	}
	return res
}

func main() {
	fmt.Println(gasStationSolution([]int{1,2,3,4,5}, []int{3,4,5,2,2}))
}

func gasStationSolution(gas , cost []int) int {
	res := -1
	for i := range gas {
		if gasStationSolution2(i, gas, cost) {
			res = i
			break
		}
	}
	return res
}

func gasStationSolution2(start int, gas, cost []int) bool {
	i, rest := start, 0
	for {
		if rest + gas[i] < cost[i] {
			fmt.Println("false", start, i)
			return false
		}
		rest = rest + gas[i] - cost[i]
		i += 1
		if i == len(gas) {
			i = 0
		}
		if i == start {
			fmt.Println("true", start, i)
			return true
		}
	}
}

