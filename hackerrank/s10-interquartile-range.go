/*
 Problem : https://www.hackerrank.com/challenges/s10-interquartile-range
 Author : Hasibul Amin Hemel
 */

package main
import (
    "fmt"
    "sort"
    "math"
)

func main() {
    var nn int
    fmt.Scanf("%d",&nn)
    xx := make([]int,nn)
    y := make([]int,nn)
    for i := 0; i < nn; i++{
        fmt.Scanf("%d",&xx[i])
    }
    for i := 0; i < nn; i++{
        fmt.Scanf("%d",&y[i])
    }
    var x []int
    
    n := 0
    
    for i := 0; i < nn; i ++{
            for k := 0; k < y[i] ; k++{
               x = append(x, xx[i])
               n++
            }
    }
    
    sort.Ints(x)
    
    
    var q1,q3 float64
    var hi float64 = 2.00
    
    if n % 2 != 0 {
        q1 = float64(x[n/4]+x[(n/4)-1])/hi
        q3 = float64(x[(n/4)+(n/2)]+x[(((n/4)+(n/2))+1)])/hi
    } else if n % 2 == 0 && (n/2) % 2 == 0 {

        q1 = float64(x[n/4]+x[(n/4)-1])/hi
        q3 = float64(x[(n/4)+(n/2)]+x[(((n/4)+(n/2))-1)])/hi
    } else {

        q1 = float64(x[n/4])
        q3 = float64(x[(n/4)+(n/2)])
    }

    
    fmt.Printf("%.1f\n",math.Abs(q1-q3))
    
}