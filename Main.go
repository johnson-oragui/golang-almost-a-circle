package main

import (
	"fmt"

	"github.com/johnson-oragui/golang-almost-a-circle/models"
)

func TaskFourMain() {
	rec1, err := models.NewRectangle(4, 6)
	CheckNilErrorValue(err)
	rec1.Display()

}

func CheckNilErrorValue(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	TaskFourMain()
}
