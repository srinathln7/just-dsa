package main 

func isAnagram(s string, t string) bool {
    
    // Key Idea: Build the freq table for every string. Increment and decrement for the two strings
    // If its a anagram, then we should all values in the map to zero

    freq := make(map[rune]int)
    for _, ch := range s {
        freq[ch]++
    }

    for _, ch := range t {
        freq[ch]--
    }

    for _, val := range freq {
        if val != 0 {
            return false 
        }
    }

    return true
}
