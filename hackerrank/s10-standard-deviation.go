/*
 Problem : https://www.hackerrank.com/challenges/s10-standard-deviation
 Author : Hasibul Amin Hemel
 */

package main
import (
    "fmt"
    "sort"
    "math"
)




func main() {
    var n int
    fmt.Scanf("%d",&n)
  
    x := make([]int,n)
    
    for i := 0; i < n; i++ {
        fmt.Scanf("%d",&x[i])

    } 
    
    p := x
    

    sort.Ints(x)

    var mean float64
    if n % 2 != 0{
        mean = float64(x[n/2])
    }else {
        mean = float64(x[n/2]+x[(n/2)-1])/float64(2) 
    }
    
    var sum float64 = 0.0
    for i := 0; i < n; i++ {
        sum += math.Pow((math.Abs(float64(p[i])-mean)),float64(2))
    }
    
    fmt.Printf("%.1f", math.Sqrt(sum/float64(n)));
       
}