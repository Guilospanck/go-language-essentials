package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sliceToReturn := [][]uint8{}
	// dy is a slice of length dy, each element of whic is a slice of dx
	for i := 0; i < dx; i++ {
		row := []uint8{}
		for j := 0; j < dy; j++ {
			value := uint8((i + j) / 2) // different fuctions to test: i*j and i^j
			row = append(row, value)
		}
		sliceToReturn = append(sliceToReturn, row)
	}

	return sliceToReturn
}

func main() {
	pic.Show(Pic)
}
