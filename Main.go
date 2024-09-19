package main

import (
	"fmt"

	"github.com/johnson-oragui/golang-almost-a-circle/models"
)

func TaskThreeMain() {
	rec1, err := models.NewRectangle(3, 2)
	CheckNilErrorValue(err)
	fmt.Println(rec1.Area())

	rec2, err := models.NewRectangle(2, 10)
	CheckNilErrorValue(err)
	fmt.Println(rec2.Area())

	rec3, err := models.NewRectangle(8, 7, 0, 0, 12)
	CheckNilErrorValue(err)
	fmt.Println(rec3.Area())
}

func CheckNilErrorValue(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	TaskThreeMain()
}
