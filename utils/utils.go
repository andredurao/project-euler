package main

func contains(m map[int]struct{}, n int) bool {
	_, isPresent := m[n]
	return isPresent
}

func appendInMap(m map[int]struct{}, n int) {
	if !contains(m, n) {
		m[n] = struct{}{}
	}
}

func keys(m map[int]struct{}) []int {
	items := make([]int, len(m))
	i := 0
	for k, _ := range m {
		items[i] = k
		i++
	}
	return items
}
