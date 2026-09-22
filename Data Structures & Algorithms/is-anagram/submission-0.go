import(
	"slices"
)


func isAnagram(s string, t string) bool {
	//Using a for loop would work in this case but run theruntime up Significantly
	//Map the individual strings to as an array and then loop comparing them
	Schar := []byte(s)
	Tchar := []byte(t)

	slices.Sort(Schar)
	slices.Sort(Tchar)

	anagram := true

	if len(Schar) != len(Tchar){
		anagram = false
	}else{
		for i := 0; i < len(Schar); i++{
			if Schar[i] != Tchar[i]{
				anagram = false
			}
		}
	}

	return anagram


}
