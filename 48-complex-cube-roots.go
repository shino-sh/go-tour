package main

import (
    "fmt"
    "math/cmplx"
)

func Cbrt(x complex128) complex128 {
    z := x
    for i:=0; i<10; i++{
      z = z - (z*z*z - x) / (3 * z*z)
    }
    return z
}

func main() {
    fmt.Println("cube root of 2 :", Cbrt(2))
    fmt.Println("cmplx.Pow      :", cmplx.Pow(Cbrt(2), 3))
}
