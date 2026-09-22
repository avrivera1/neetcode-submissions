func groupAnagrams(strs []string) [][]string {
	//I want to put the anagrams in a hash table but I wont be able to split them afterwards

	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++ // Assumes lowercase a-z
		}
		// Go allows arrays as map keys because they are comparable
		groups[count] = append(groups[count], s)
	}

	// Flatten the map values into the final result
	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result

}