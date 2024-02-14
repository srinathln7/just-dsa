package main

import "strconv"

func calPoints(operations []string) int {
	var result, n int
	var record []int

	for _, str := range operations {
		n = len(record)
		switch str {
		case "+":
			if n >= 2 {
				record = append(record, (record[n-1] + record[n-2]))
			}
		case "D":
			if n >= 1 {
				record = append(record, 2*record[n-1])
			}
		case "C":
			if n >= 1 {
				record = record[:(n - 1)]
			}
		default:
			num, _ := strconv.Atoi(str)
			record = append(record, num)
		}
	}

	for _, val := range record {
		result += val
	}

	return result
}
