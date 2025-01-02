package main

import (
	"go_study/usePkg/custompkg"
	"go_study/usePkg/exinit"
)

func main() {
	custompkg.CustomFunc()
	custompkg.CustomFunc2()

	exinit.PrintD()

	// S := custompkg.Student{}

	// fmt.Printf("%+v\n", S)
	// data := []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 3, 0.3, 15.5, 1, 6}
	// graph := asciigraph.Plot(data)
	// fmt.Println(graph)
}
