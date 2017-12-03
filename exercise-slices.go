package main

import (
	"golang.org/x/tour/pic"
)

/**
 * https://tour.golang.org/moretypes/18
 */

func Pic(dx, dy int) [][]uint8 {
    // makeSlice of type uint8
	ret := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			ret[i][j] = uint8(i^j+(i+j)/2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
