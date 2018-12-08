package diffs

func DiffSetSlices(expected, actual []string) (excess, missing []string) {
	expectedSet := make(map[string]bool)
	for _, el := range expected {
		expectedSet[el] = true
	}
	actualSet := make(map[string]bool)
	for _, el := range actual {
		actualSet[el] = true
	}
	for el, _ := range expectedSet {
		if _, ok := actualSet[el]; !ok {
			missing = append(missing, el)
		}
	}
	for el, _ := range actualSet {
		if _, ok := expectedSet[el]; !ok {
			excess = append(excess, el)
		}
	}
	return
}
