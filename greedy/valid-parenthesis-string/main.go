package main

func checkValidString(s string) bool {

	// Key Idea: As we range over the characters, we keep track of the number of unmatached left parenthesis.
	// When we have a wild card character, we can make three decisions. We ignore the empty. If we choose to
	// put `(` then no. of unmatched left parenthesis will increase and `)` then no of unmatched left parenthesis
	// will decrease.

	// leftMin : Represents the value when choosing wild card `*` as `)`
	// leftMax : Represents the value when choosing wild card `*` as `(`
	var leftMin, leftMax int
	for _, ch := range s {
		switch ch {
		case '(':
			leftMin, leftMax = leftMin+1, leftMax+1
		case ')':
			leftMin, leftMax = leftMin-1, leftMax-1
		default:
			leftMin, leftMax = leftMin-1, leftMax+1
		}

		if leftMax < 0 {
			return false
		}

		// An intuitive explanation: As we progress through the string, our minimum and maximum counts of unmatched left parentheses (`leftmin` and `leftmax`)
		// dynamically change. If the `leftmin` becomes negative, it indicates that we've encountered more right parentheses than the total number of corresponding
		// left parentheses and asterisks seen so far. In such cases, we can revise the previous characters to include an empty space, utilizing the wildcard '*' as
		// an optional left parenthesis. This gives the string another chance to remain valid. However, if the `leftmax` becomes negative, it signifies an irrecoverable
		// situation. This occurs when, despite using all wildcards as left parentheses, the count of right parentheses exceeds the count of remaining unmatched left
		// parentheses and asterisks. In essence, it means that the string cannot be
		// balanced, rendering it invalid. This approach ensures that the string's validity is continuously monitored and maintained throughout the traversal.

		// Rest leftMin to zero if it ever becomes negative
		leftMin = max(0, leftMin)
	}

	return leftMin == 0
}
