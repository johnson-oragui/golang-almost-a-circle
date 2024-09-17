package models

/*
Struct Rectangle:

Define a struct named Rectangle that embeds the Base struct.
Define public fields: width, height, x, and y of type int.
NewRectangle function:

Implement a function named NewRectangle that takes a variable number of int arguments.
Inside the function:
Create a new instance of the Rectangle struct using &Rectangle{}.
Extract values for width, height, x, and y from the variadic arguments.
Create a new Base instance using NewBase and assign it to the embedded Base field of the Rectangle instance.
Return a pointer to the newly created Rectangle struct.
*/

type Rectangle struct {
	width int
	height int
	x int
	y int
	Base
}


func NewRectangle(values...int) *Rectangle {
	rectangle := &Rectangle{
		width: 0,
		height: 0,
		x: 0,
		y: 0,
	}

	switch len(values) {
	case 0:
		rectangle.Base = *NewBase()
		return rectangle
	case 1:
		rectangle.width = values[0]
	case 2:
		rectangle.width = values[0]
		rectangle.height = values[1]
	case 3:
		rectangle.width = values[0]
		rectangle.height = values[1]
		rectangle.x = values[2]
	case 4:
		rectangle.width = values[0]
		rectangle.height = values[1]
		rectangle.x = values[2]
		rectangle.y = values[3]
	default:
		rectangle.width = values[0]
		rectangle.height = values[1]
		rectangle.x = values[2]
		rectangle.y = values[3]
		rectangle.Base = *NewBase(values[4])
		return rectangle
	}

	rectangle.Base = *NewBase()
	return rectangle
}
