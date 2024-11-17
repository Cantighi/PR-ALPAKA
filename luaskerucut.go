package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Menghitung Luas Kerucut")
	fmt.Println("diketahui kerucut dengan tinggi 20 dan jari jari 5")
	hitungluaskerucut()

}

/*yang dibutuhkan tinggi miring, luas alas, luas selimut*/
/*yang diketahui tinggi (t) = 20 dan jari jari (r) = 5*/

func hitungluaskerucut() {
	var r float64
	var t float64

	r = 5
	t = 20

	tinggimiring := math.Sqrt(r*r + t*t)
	luasalas := math.Pi * r * r
	luasselimut := math.Pi * r * tinggimiring
	luaskerucut := luasalas + luasselimut

	fmt.Println("tinggi miring adalah", tinggimiring)
	fmt.Println("luas alas adalah", luasalas)
	fmt.Println("luas selimut adalah", luasselimut)
	fmt.Println("luas kerucut adalah", luaskerucut)

}
