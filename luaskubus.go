package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calculating the surface area of a cube")
	countcubesurfacearea()
	customcubesurface()
}

/* the side is 5*/
func countcubesurfacearea() {
	var side int
	side = 5
	surfacearea := 6 * side * side

	fmt.Println("the surface area of the cube is:", surfacearea)

}

/*the side can be costumed*/

func customcubesurface() {
	var customside float32
	fmt.Print("Input the side of the cube: ")
	fmt.Scan(&customside)
	customcubesurfacearea := 6 * customside * customside

	fmt.Println("The surface of the cube with side", customside, "is:", customcubesurfacearea)
}
