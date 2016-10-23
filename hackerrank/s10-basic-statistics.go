package main
import (
    "fmt"
    "sort"
)

func tSum (n []int) float64{
    ret := 0
    for i := 0; i < len(n); i++{
        ret += n[i]
    }
    return float64(ret);
}




func main() {
    
    var n int
    fmt.Scanf("%d", &n)
    num  := make([]int,n)
    mp  := make(map[int]int)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &num[i])
        mp[num[i]]++
    }
 
    sort.Ints(num)
    
   
    fmt.Printf("%.1f\n",tSum(num)/float64(n))
    var median float64
    if n % 2 != 0 {
        median = float64(num[n/2])
    }else {
        median = float64(float64(num[n/2]+num[(n/2)-1])/float64(2))
    }
    
    fmt.Printf("%.1f\n",median)
    
    var keys []int
    for k := range mp {
        keys = append(keys, k)
    }
    
    sort.Sort(sort.Reverse(sort.IntSlice(keys)))
 
    min := -1
    mode := -1
    for _, k := range keys {
        if mp[k] >= min {
           min = mp[k]
           mode = k
        } 
    }
    
    fmt.Println(mode)
       
}