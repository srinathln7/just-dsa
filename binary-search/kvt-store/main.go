package main 

type Element struct {
    val string
    timeStamp int
}

type TimeMap struct {
    kvt map[string][]Element
}


func Constructor() TimeMap {
    return TimeMap{kvt : make(map[string][]Element)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.kvt[key] = append(this.kvt[key], Element{val: value, timeStamp: timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    var res string    
    if _, exists := this.kvt[key]; !exists {
        return res   
    }


    elements := this.kvt[key] 
    l, r := 0, len(elements) - 1
    for l <= r {
        mid := l + (r-l)/2
        val, midTS := elements[mid].val,  elements[mid].timeStamp 
        if midTS == timestamp {
            return val
        }

        // IMPORTANT: For ex [1, 2, 4,  6,  8, 10, 12] => O/P (7 -> 6), ( 5-> 4), ( 3 -> 2)
        // If key exists, to find the element closest to the timestamp, we can run a binary search
        // where we compare the mid-value and check if the mid-value is lesser than the timestamp 
        // If yes, then we try to move our left ptr to find if we can get a closer value to the timestamp
        // If not, we move our right ptr to find the closest timestamp
        if midTS < timestamp {
            res = val  // Capture this value after every iteration
            l = mid+1
        } else {
            r = mid-1
        }
    }

   return res
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
