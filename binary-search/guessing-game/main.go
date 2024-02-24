package main 

/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 1, n

    var try int 
    for l <= r {
        try = l + (r-l)/2
        if guess(try) == 0 {
            return try
        } 

        if guess(try) < 0 {
            r = try-1
        } else {
            l = try+1 
        }
    }

 return -1    
}
