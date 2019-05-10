package ly

import "strings"

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, e := range strings.Fields(s) {
		for _, i := range e {
			result[string(i)] += 1
		}
	}
	return result
}
