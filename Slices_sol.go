package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var res [][]uint8
	for i := range make([]int, dx) {
		res = append(res, make([]uint8, dy))
		for j := range make([]int, dy) {
			res[i][j] = uint8((i + j) / 2)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
