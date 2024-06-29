package set

// Diff A - Bの差集合を求める
func Diff(srcSet, ignoreSet map[string]struct{}) []string {
	var result []string
	for file := range srcSet {
		if _, found := ignoreSet[file]; !found {
			result = append(result, file)
		}
	}
	return result
}

// SliceToSet converts a slice of strings to a set (map[string]struct{}).
func SliceToSet(slice []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, item := range slice {
		set[item] = struct{}{}
	}
	return set
}
