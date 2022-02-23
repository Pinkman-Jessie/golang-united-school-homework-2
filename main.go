package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	//CalcSquare(5.0, 4)
}
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0
type customtype int
func CalcSquare(sideLen float64, sidesNum customtype) (res float64) {

	if sidesNum == SidesTriangle {

		res = sideLen * sideLen
	} else if sidesNum == SidesSquare {
		res = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		res = math.Pi * (sideLen * sideLen)
	} else {
		res = 0
	}
	//fmt.Println("yooo",res)
	return
	
}
