package main

import "math"


type Shape interface{
    Area() float64
}

type Circle struct{
    Radius float64
}

func (circle Circle) Area() float64{

    return circle.Radius * circle.Radius * math.Pi
} 



type Rectangle struct{
    Width float64
    Height float64
}


func (rectangle Rectangle) Area() float64{
    return rectangle.Width * rectangle.Height
}


func  (rectangle Rectangle) Perimeter() float64{
    return 2* (rectangle.Width + rectangle.Height)
}


type Triangle struct{
    Height float64
    Base float64
}

func (triangle Triangle) Area() float64{
    return (triangle.Base * triangle.Height)/2 
}




