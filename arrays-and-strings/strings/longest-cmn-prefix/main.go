package main

func longestCommonPrefix(strs []string) string {

	prefix := strs[0]
	for i, str := range strs {
		if i == 0 {
			continue
		}

		j := 0
		for j < len(prefix) && j < len(str) && str[j] == prefix[j] {
			j++
		}
		prefix = prefix[:j]
	}

	return prefix
}
