package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Calculating area of a circle")
	countcirclearea()
	customcirclearea()

}

/*r = 5*/
func countcirclearea() {
	var r float64
	r = 5
	circlearea := math.Pi * r * r

	fmt.Println("Area of a circle if r = 5 is:", circlearea)

}

/*the radius can be customized*/
func customcirclearea() {
	var customradius float64
	fmt.Print("Input the radius of the circle")
	fmt.Scan(&customradius)
	customcircleradius := math.Pi * customradius * customradius

	fmt.Println("Area of a circle if the radius is", customradius, "is", customcircleradius)
}
