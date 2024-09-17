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

func main() {
	TaskZeroMain()
	
}
