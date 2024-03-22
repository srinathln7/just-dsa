package main 

func groupAnagrams(strs []string) [][]string {
    
    // Key: Lets build a hashmap and everytime we insert an element, lets sort it and check if it already exists

    n := len(strs)
    result := make([][]string, n)

    idxMap := make(map[string]int)
    for i, str := range strs {
        sortedStr := sortString(str)
        if idx, exists := idxMap[sortedStr]; exists {
            result[idx] = append(result[idx], str)
        } else {
            idxMap[sortedStr] = i
            result[i] = []string{str}
        }
    }  

    var res [][]string
    for _, arr := range result {
        if len(arr) != 0 {
            res = append(res, arr)
        }
    }

    return res
}

func sortString(str string) string {
    s := strings.Split(str, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
