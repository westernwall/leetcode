package TwoSum

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums{
		if index, ok := m[v]; ok {
			return []int{index, i}
		}

		m[target - v] = i
	}

	return nil
}