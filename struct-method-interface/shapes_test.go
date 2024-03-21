package main


import "testing"




func TestPerimeter(t *testing.T){
    rectangle := Rectangle{10.0,10.0}
    got := rectangle.Perimeter()
    want := 40.0
    if got != want{
        t.Errorf("got %g wanted %g", got, want)
    }
}

func TestArea(t *testing.T){
    
    areaTests := []struct{
        name string
        shape Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{10.0,10.0},hasArea: 100},
        {name: "Circle",shape: Circle{10.0},hasArea: 314.1592653589793},
        {name: "Triangle",shape: Triangle{2.0,4.0}, hasArea: 4.0},
    }
    
    for _, tt := range areaTests{
        got := tt.shape.Area()
        if got  != tt.hasArea{
            t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
        }
    }


}





