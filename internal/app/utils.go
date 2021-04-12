package app

import (
	"fmt"
	"sort"
)

func uniqueStringSlice(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if entry == "" {
			continue
		}
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func createConcatenatedString(data []string) string {
	var result string
	if len(data) > 0 {
		data = uniqueStringSlice(data)
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		for i, v := range data {
			if i > 0 {
				result = fmt.Sprintf("%s,%s", result, v)
			} else {
				result = v
			}
		}
	}
	return result
}
