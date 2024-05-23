package main

func maxNumberOfBalloons(text string) int {

	// Key Idea: Count the freq of individual characters and take the min. of the all characters that appear in the desired
	// word "balloon". Account for characters that appear twice.

	m, n := len("balloon"), len(text)
	if n < m {
		return 0
	}

	count := [26]int{}
	for i := 0; i < n; i++ {
		count[text[i]-'a']++
	}

	// We need 2 l's and 2 o's to form the word "balloon". Hence, we divide those count instances by 2.
	return min(count['b'-'a'], count['a'-'a'], count['l'-'a']/2, count['o'-'a']/2, count['n'-'a'])
}
