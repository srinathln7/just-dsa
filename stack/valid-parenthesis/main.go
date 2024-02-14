package main

func isValid(s string) bool {
    var stack []string
    var n int
    var ch  string

    for _, str := range s {
        n = len(stack)
        ch = string(str)
        switch ch {
            case "(":
              stack = append(stack, ")")
            case "[":
              stack = append(stack, "]")
             case "{":
              stack = append(stack, "}")
            default:
                 if n == 0 || stack[n-1] != ch {
                     return false
                 } else {
                     stack = stack[:n-1]
                 }    
        }
    }

    return len(stack)==0
}


