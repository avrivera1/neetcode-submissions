func twoSum(nums []int, target int) []int {
	//replace nested for loop by storing visited objects in a hash an revist it
	//Counting difference
    seen := make(map[int]int)


	for i, currentNum := range nums{

		difference := target - currentNum

		if index, found := seen[difference]; found {
			// Found the pair! Return their indices
			return []int{index, i}
		}

		seen[currentNum] = i
	}
	
	return nil


}
