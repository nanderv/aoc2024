package chal5

// numsFulfillRules takes a manual (both as list ánd map since we have both available, and the set of rules that has to be fulfilled
// it returns three values:
// - bool: does the list of numbers fulfil all of the rules?
// - two integers: if the rules are failed, it returns two positions in the manual of elements that are not in the right order.
func manualFulfillsRules(nums []int, numPositions map[int]int, rules map[int][]int) (bool, int, int) {
	for k, v := range nums {
		// Number isn't the left part of any rule (so doesn't have any numbers that have to be right of it)
		rlz, ok := rules[v]
		if !ok {
			continue
		}
		for _, r := range rlz {
			if l, ok := numPositions[r]; ok {
				if l < k {
					return false, k, l
				}
			}
		}
	}
	return true, 0, 0
}
