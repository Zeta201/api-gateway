package main

import (
	"fmt"
	"math"
)

// type demoInterface interface{
// 	add(int, int) int
// }

// type demoInterface2 interface{
// 	add(int ,int) int

// }

type Point struct{
	x int
	y int
}

type Triangle struct{
	points []Point
}

func (t *Triangle) Area() float64 {
	return 1.5
}

func (p *Point) Length() int{
	return int(math.Abs(float64(p.x-p.y)))
}



func initBytes(byteSlice *[]byte){
	var name string = "abcd"
	*byteSlice = append(*byteSlice, []byte(name)...)
	fmt.Println(*byteSlice)
}
func myFunc(data interface {}){
	fmt.Println(data)
	p := data.(Point)
	fmt.Println(p.Length())


}

const (
	_ int = iota
	x 
	y 
	Z
	t
)

func init(){
	fmt.Println("Initializing")
	fmt.Println("--------------------------------")
}
func main(){


	t := Triangle{
		points: []Point{
			Point{1,2},
            Point{3,4},
           Point {5,6},
		},
	}
	

	fmt.Println(t.Area())
	fmt.Println(t)
	var name = []byte{}
	initBytes(&name)
	fmt.Println(name)

	// temp := x
	// fmt.Println(temp)



	// p := Point{
	// 	x:10,
	// 	y:20,
	// }
	// fmt.Println(p.Length())
	// myFunc(p)

	// // match, _ := regexp.MatchString("p[a-z]+ch","peach")
	// // fmt.Println(match)


	// r, _ := regexp.Compile("p[A-Z]+[a-z]+ch")
	// fmt.Println(r.String())
	// fmt.Println(r.MatchString("peach"))
	// fmt.Printf("%T",r)
	
	
}