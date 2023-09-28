package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
	sides int
}

func (r rectangle) area() float64  { //A interface shape esta sendo implementada pelo objeto rectangle
	return r.height * r.width
}

func (r rectangle) perimeter() float64  { //A interface shape esta sendo implementada pelo objeto rectangle
	return 2*r.width + 2*r.height
}


func Calculate(s shape)  {
	fmt.Print(s.area())
	fmt.Print(" ")
	fmt.Print(s.perimeter())
}

func main()  {
	rec := rectangle{
		width: 5,
		height:8,
		sides: 4,
	}
	Calculate(rec) //por mais que o tipo de calculate é um shape ele aceita o objeto rectangle, pois as funçoes de shape aceita 
}