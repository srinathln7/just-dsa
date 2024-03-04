package main

type tracker struct {
	firstIndex int
	count      int
}

func firstUniqChar(s string) int {
	trackerMap := make(map[rune]tracker)

	for index, char := range s {
		if val, exists := trackerMap[char]; exists {
			newcount := val.count + 1
			trackerMap[char] = tracker{count: newcount}
			continue
		}

		trackerMap[char] = tracker{firstIndex: index, count: 1}
	}

	var resultIndex int = -1
	for _, val := range trackerMap {
		if val.count == 1 {
			if resultIndex == -1 {
				resultIndex = val.firstIndex
			} else {
				if val.firstIndex < resultIndex {
					resultIndex = val.firstIndex
				}
			}

		}
	}

	return resultIndex
}
