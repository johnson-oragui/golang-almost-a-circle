package models

/*
Struct Rectangle:

Define a struct named Rectangle that embeds the Base struct.
Define public fields: Width, Height, x, and y of type int.
NewRectangle function:

Implement a function named NewRectangle that takes a variable number of int arguments.
Inside the function:
Create a new instance of the Rectangle struct using &Rectangle{}.
Extract values for Width, Height, x, and y from the variadic arguments.
Create a new Base instance using NewBase and assign it to the embedded Base field of the Rectangle instance.
Return a pointer to the newly created Rectangle struct.
*/

type Rectangle struct {
	Width int
	Height int
	X int
	Y int
	Base
}


func NewRectangle(values...int) *Rectangle {
	rectangle := &Rectangle{
		Width: 0,
		Height: 0,
		X: 0,
		Y: 0,
	}

	switch len(values) {
	case 0:
		rectangle.Base = *NewBase()
		return rectangle
	case 1:
		rectangle.Width = values[0]
	case 2:
		rectangle.Width = values[0]
		rectangle.Height = values[1]
	case 3:
		rectangle.Width = values[0]
		rectangle.Height = values[1]
		rectangle.X = values[2]
	case 4:
		rectangle.Width = values[0]
		rectangle.Height = values[1]
		rectangle.X = values[2]
		rectangle.Y = values[3]
	default:
		rectangle.Width = values[0]
		rectangle.Height = values[1]
		rectangle.X = values[2]
		rectangle.Y = values[3]
		rectangle.Base = *NewBase(values[4])
		return rectangle
	}

	rectangle.Base = *NewBase()
	return rectangle
}
