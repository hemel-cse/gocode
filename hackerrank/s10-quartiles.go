/*
 Problem : https://www.hackerrank.com/challenges/s10-quartiles
 Author : Hasibul Amin Hemel
 */

package main
import (
    "fmt"
    "sort"
)

func main() {
    var n int
    fmt.Scanf("%d",&n)
    x := make([]int,n)
    for i := 0; i < n; i++{
        fmt.Scanf("%d",&x[i])
    }
    sort.Ints(x)
    
    var q1,q2,q3 int
    
    if n % 2 != 0 {
        q2  = x[n/2]
        q1 = (x[n/4]+x[(n/4)-1])/2
        q3 = (x[(n/4)+(n/2)]+x[(((n/4)+(n/2))+1)])/2
     } else {
        q2 = (x[n/2]+x[(n/2)-1])/2
        q1 = (x[n/4])/2
        q3 = x[(n/4)+(n/2)]/2
    }
    
    fmt.Printf("%d\n%d\n%d\n",q1,q2,q3)
    
}