package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := x
    for i:=0; i<10; i++{
      z = z - ( (math.Pow(z, 2) - x) / (2 * z) )
    }
    return z
}

func main() {
    for i:=0; i<10; i++{
        x := float64(i)
      fmt.Println(i, Sqrt(x), math.Sqrt(x))
    }
}
