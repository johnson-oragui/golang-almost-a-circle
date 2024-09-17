package models

/*
Create a package named models with the following functionality:

Struct Base:

Define a struct named Base with a single field Id of type int.
NewBase function:

Implement a function named NewBase that takes an optional int argument id.
Inside the function:
Create a new instance of the Base struct using &Base{}.
If id is provided and greater than 0, assign it to the Id field of the new instance.
Otherwise, increment a global variable named nbObjects.
Assign the incremented value of nbObjects to the Id field of the new instance.
Return a pointer to the newly created Base struct.
*/
type Base struct {
	Id int
}

var nbObjects int = 0

func NewBase(id ...int) *Base {
	// create a new instance of Base
	base := &Base{}

	// If an id is provided, assign it to the instance's Id
	if len(id) > 0 && id[0] > 0 {
		base.Id = id[0]
	} else {
		// Otherwise, increment the nbObjects and assign the new value to Id
		nbObjects++
		base.Id = nbObjects
	}
	return base
}
