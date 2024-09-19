package main

import (
	"fmt"

	"github.com/johnson-oragui/golang-almost-a-circle/models"
)

func TaskTwoMain() {
	if r1, err := models.NewRectangle(10, 2); err != nil {
		fmt.Printf("%v\n", err)
	} else {
	fmt.Printf("width: %d, height: %d, x: %d, y: %d, Id: %d\n", r1.Width, r1.Height, r1.X, r1.Y, r1.Id)
	}

	if r2, err := models.NewRectangle(12, 0); err != nil {
		fmt.Printf("%v\n", err)
	} else {
	fmt.Printf("width: %d, height: %d, x: %d, y: %d, Id: %d\n", r2.Width, r2.Height, r2.X, r2.Y, r2.Id)
	}

	if r3, err := models.NewRectangle(10, 2, 0, 0, 12); err != nil {
		fmt.Printf("%v\n", err)
	} else {
	fmt.Printf("width: %d, height: %d, x: %d, y: %d, Id: %d\n", r3.Width, r3.Height, r3.X, r3.Y, r3.Id)
	}

	if r4, err := models.NewRectangle(10); err != nil {
		fmt.Printf("%v\n", err)
	} else {
	fmt.Printf("width: %d, height: %d, x: %d, y: %d, Id: %d\n", r4.Width, r4.Height, r4.X, r4.Y, r4.Id)
	}
	if r5, err := models.NewRectangle(20, 10, -1, -2); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("width: %d, height: %d, x: %d, y: %d, Id: %d\n", r5.Width, r5.Height, r5.X, r5.Y, r5.Id)
	}

}

func main() {
	TaskTwoMain()
	
}
