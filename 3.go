package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func didalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) < float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1):")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2):")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&x, &y)

	inCircle1 := didalam(cx1, cy1, r1, x, y)
	inCircle2 := didalam(cx2, cy2, r2, x, y)

	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
