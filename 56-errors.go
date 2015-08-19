package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(f float64) (float64, error) {
    if f <= 0 {
      return 0, ErrNegativeSqrt(f)
    }
    
    const delta = 1e-10
    z := f
    for {
        n := z - (z*z - f) / (2 * z)
        if math.Abs(n-z) < delta { break }
        z = n
    }
    return z, nil
}

/*
func SqrtPrint(val float64, err error) {
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(val)
    }
}
*/

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
    
    //SqrtPrint(Sqrt(2))
    //SqrtPrint(Sqrt(-2))
}
