package main

func myPowAlt(x float64, n int) float64 {
    
    if x == 0.0 {
        return 0.0
    }

    if n == 0 || x == 1.0 {
        return 1.0
    }


    var value float64 = 1.0
    for i :=1; i <= abs(n); i++ {
        value *= x
    }

    if n < 0 {
        return 1/value
    }

   return value 
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
