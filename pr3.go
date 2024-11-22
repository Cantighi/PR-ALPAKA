package main

import (
	"fmt"
	"math"
)

// cube surface//
func cubesurface(side float64) float64 {
	return math.Pow(side, 2) * 6
}

// circle area//
func circlearea(radius float64) float64 {
	return math.Pi * radius * radius
}

// slainheightcone//
func slainheight(height, radius float64) float64 {
	return math.Sqrt(height*height + radius*radius)
}

// cone area//
func conearea(height, radius float64) float64 {
	slaintheight := math.Sqrt(radius*radius + height*height)
	return math.Pi * radius * (radius + slaintheight)
}

// print//
func main() {
	var choice int
	fmt.Println("count the area of: ")
	fmt.Println("1. Surface of a cube")
	fmt.Println("2. Surface of a circle")
	fmt.Println("3. surface of a cone")
	fmt.Scan(&choice)

	//pick one of the func//

	if choice == 1 {
		var side float64
		fmt.Println("side of the cube:")
		fmt.Scan(&side)
		area := cubesurface(side)
		fmt.Println("the area of a cube if side:", &side, area)
	}
	if choice == 2 {
		var radius float64
		fmt.Println("radius of the circle:")
		fmt.Scan(&radius)
		area := circlearea(radius)
		fmt.Println("the area of a circle if side:", &radius, area)
	}
	if choice == 3 {
		var radius float64
		var height float64
		fmt.Println("radius of the cone:")
		fmt.Scan(&radius)
		fmt.Println("height of the cone:")
		fmt.Scan(&height)
		slainheight := slainheight(height, radius)
		fmt.Println("the slaint height of the cone:", slainheight)
		conearea := conearea(radius, height)
		fmt.Println("the area of the cone:", conearea)
	}
}
