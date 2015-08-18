package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    image := make([][]uint8, dy)
    for j := range image {
        image[j] = make([]uint8, dx)
        for i := range image[j] {
          image[j][i] = uint8(i ^ j)
        }
    }
    return image
}

func main() {
    pic.Show(Pic)
}
