package main

import (
	"fmt"

	"github.com/johnson-oragui/golang-almost-a-circle/models"
)

func TaskZeroMain() {
	b1 := models.NewBase()

	b2 := models.NewBase()

	b3 := models.NewBase()

	b4 := models.NewBase(12)

	b5 := models.NewBase()

	fmt.Printf("%d\n%d\n%d\n%d\n%d\n", b1.Id, b2.Id, b3.Id, b4.Id, b5.Id)
}

func TaskOneMain() {
	r1 := models.NewRectangle(10, 2)

	fmt.Println(r1.Id)
	r2 := models.NewRectangle(2, 10)

	fmt.Println(r2.Id)

	r3 := models.NewRectangle(10, 2, 0, 0, 12)

	fmt.Println(r3.Id)

	r4 := models.NewRectangle()

	fmt.Println(r4.Id)
}

func main() {
	// TaskZeroMain()
	TaskOneMain()
	
}
