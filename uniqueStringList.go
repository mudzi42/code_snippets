package codesnippets


// uniqueStringList

func uniqueStringList(list []string) []string {
	u := make([]string, 0, len(list))
	m := make(map[string]bool)

	for _, val := range list {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}