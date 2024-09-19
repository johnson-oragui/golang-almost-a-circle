package main

import (
	"fmt"

	"github.com/johnson-oragui/golang-almost-a-circle/models"
)

func TaskSixMain() {
	rec1, err := models.NewRectangle(4, 6, 2, 1, 12)
	CheckNilErrorValue(err)
	fmt.Println(rec1)

	rec2, err := models.NewRectangle(5, 5, 1)
	CheckNilErrorValue(err)
	fmt.Println(rec2)

}

func CheckNilErrorValue(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	TaskSixMain()
}
