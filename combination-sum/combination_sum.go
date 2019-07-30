package combinationsum

import "sort"

func combinationSum(candidates []int, target int) (out [][]int) {
	out = [][]int{}

	var findCombos func([]int, []int, int)
	findCombos = func(combination []int, currentCandidates []int, currentTarget int) {
		if currentTarget == 0 {
			out = append(out, combination)
		} else {
			for k, v := range currentCandidates {
				if v <= currentTarget {
					cpy := []int{}
					cpy = append(cpy, combination...)
					cpy = append(cpy, v)
					findCombos(cpy, currentCandidates[k:], currentTarget-v)
				} else {
					break
				}
			}
		}
	}
	sort.Ints(candidates[:])
	findCombos([]int{}, candidates, target)
	return
}
