package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	x int
	y int
}

func CreatePoint(x int, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (v *Point) GetDistance(point Point) float64 {
	xWidth := v.x - point.x
	yWidth := v.y - point.y

	return math.Sqrt(float64(xWidth*xWidth + yWidth*yWidth))
}

func main() {
	var x1 int
	fmt.Print("Введите x1: ")
	_, err := fmt.Fscan(os.Stdin, &x1)
	if err != nil {
		panic(err)
	}

	var y1 int
	fmt.Print("Введите y1: ")
	_, err = fmt.Fscan(os.Stdin, &y1)
	if err != nil {
		panic(err)
	}

	var x2 int
	fmt.Print("Введите x1: ")
	_, err = fmt.Fscan(os.Stdin, &x2)
	if err != nil {
		panic(err)
	}

	var y2 int
	fmt.Print("Введите y2: ")
	_, err = fmt.Fscan(os.Stdin, &y2)
	if err != nil {
		panic(err)
	}

	p1 := CreatePoint(x1, y1)
	p2 := CreatePoint(x2, y2)

	fmt.Printf("Расстояние между точками = %v\n", p1.GetDistance(p2))
}
