package models

import (
	"errors"
	"fmt"
)


type Rectangle struct {
	Width int
	Height int
	X int
	Y int
	Base
}

func (r *Rectangle) setWidth(width int) error {
	if width <= 0 {
		return errors.New("width must be > 0")
	}
	r.Width = width
	return nil
}

func (r *Rectangle) setHeight(height int) error {
	if height <= 0 {
		return errors.New("height must be > 0")
	}
	r.Height = height
	return nil
}

func (r *Rectangle) setX(x int) error {
	if x < 0 {
		return errors.New("x must be >= 0")
	}
	r.X = x
	return nil
}

func (r *Rectangle) setY(y int) error {
	if y < 0 {
		return errors.New("y must be >= 0")
	}
	r.Y = y
	return nil
}

func checkRectangleValues(rectangle *Rectangle, values ...int) error {
	if len(values) < 2 {
		return errors.New("width and height must be set and greather than zero")
	}

	err := rectangle.setWidth(values[0])
	if err != nil {
		return err
	}

	if err := rectangle.setHeight(values[1]); err != nil {
		return err
	}

	if len(values) > 2 {
		if err := rectangle.setX(values[2]); err != nil {
			return err
		}
	}
	if len(values) > 3 {
		if err := rectangle.setY(values[3]); err != nil {
			return err
		}
	}
	if len(values) > 4 {
		rectangle.Base = *NewBase(values[4])
	} else{
		rectangle.Base = *NewBase()
	}
	return nil
}

func NewRectangle(values ...int) (*Rectangle, error) {
	rectangle := &Rectangle{
		Width: 0,
		Height: 0,
		X: 0,
		Y: 0,
	}

	if err := checkRectangleValues(rectangle, values...); err != nil {
		return nil, err
	}
	return rectangle, nil
}

func (r *Rectangle) Area() int {
	return r.Height * r.Width
}

func (r *Rectangle) Display() {
	for i := 1; i <= r.Height; i++ {
		for j := 1; j <= r.Width; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("[Rectangle] (%d) %d/%d - %d/%d", r.Id, r.X, r.Y, r.Width, r.Height)
}
