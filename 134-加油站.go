package main

func canCompleteCircuit(gas []int, cost []int) int {
	cur, total, start := 0, 0, 0
	for index := 0; index < len(gas); index++ {
		cur += gas[index] - cost[index]
		total += gas[index] - cost[index]
		if cur < 0 {
			cur = 0
			start = index + 1
		}
	}
	if total >= 0 {
		return start
	} else {
		return -1
	}
}
