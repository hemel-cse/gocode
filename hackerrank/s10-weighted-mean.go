/*
 Problem : https://www.hackerrank.com/challenges/s10-weighted-mean
 Author : Hasibul Amin Hemel
 */

package main
import "fmt"

func wght (n int, x []int, w []int) float64{
    sum := 0
    wht := 0
    for i := 0; i < n; i++ {
        sum += (x[i]*w[i])
        wht += w[i]
    }
    
    return float64(float64(sum) /float64(wht))
}


func main() {
    var n int
    fmt.Scanf("%d", &n)
    x := make([]int,n)
    w := make([]int,n)
    for i := 0; i < n; i++{
        fmt.Scanf("%d", &x[i])
    }
    for i := 0; i < n; i++{
        fmt.Scanf("%d", &w[i])
    }
    
    fmt.Printf("%.1f",wght(n,x,w))
    
}