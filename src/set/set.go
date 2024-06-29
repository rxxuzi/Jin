package set

// Diff computes the set difference A - B, which is the elements in A that are not in B.
func Diff(a, b []string) []string {
	var result []string
	set := SliceToSet(b)

	for _, item := range a {
		if _, found := set[item]; !found {
			result = append(result, item)
		}
	}

	return result
}

// And computes the intersection of sets A and B, which is the elements common to both A and B.
func And(a, b []string) []string {
	var result []string
	set := SliceToSet(a)

	for _, item := range b {
		if _, found := set[item]; found {
			result = append(result, item)
		}
	}

	return result
}

// Or computes the union of sets A and B, which is all elements that are in either A or B.
func Or(a, b []string) []string {
	result := make(map[string]struct{})
	var finalResult []string
	for _, item := range a {
		result[item] = struct{}{}
	}
	for _, item := range b {
		result[item] = struct{}{}
	}

	for item := range result {
		finalResult = append(finalResult, item)
	}

	return finalResult
}

// SliceToSet converts a slice of strings to a set (map[string]struct{}).
func SliceToSet(slice []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, item := range slice {
		set[item] = struct{}{}
	}
	return set
}
