func twoSum(nums []int, target int) []int {
	//replace nested for loop by storing visited objects in a hash an revist it
	//Counting difference
    seen := make(map[int]int)


	for i, currentNum := range nums{

		difference := target - currentNum

		if j, found := seen[difference]; found {
			// Found the pair! Return their indices
			return []int{j, i}
		}

		seen[currentNum] = i
	}
	
	return nil


}
