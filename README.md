### GOLANG-ALMOST-A-CIRCLE


### Tasks
## Task 0.
**If it's not tested it doesn't work**
mandatory
All your files, classes and methods must be unit tested.

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go test
 ./models

ok      github.com/johnson-oragui/golang-almost-a-circle/models (cached)

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ 

File: models/models_test.go
   
## Task 1.
**Base class**
mandatory
Create a package named models with the following functionality:

Struct Base:

- Define a struct named Base with a single field Id of type int.
NewBase function:

 - Implement a function named NewBase that takes an optional int argument id.
Inside the function:
 - Create a new instance of the Base struct using &Base{}.
 - If id is provided and greater than 0, assign it to the Id field of the new instance.
 - Otherwise, increment a global variable named nbObjects.
 - Assign the incremented value of nbObjects to the Id field of the new instance.
 - Return a pointer to the newly created Base struct.

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$  cat Main.go
```
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

```

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go run Main.go

1

2

3

12

4

Github Repo: golang-almost_a_circle
File: models/base.go, models/base_test.go

## Task 2
**First Rectangle**

- Define a struct named Rectangle that embeds the Base struct.
- Define public fields: width, height, x, and y of type int.
NewRectangle function:

- Implement a function named NewRectangle that takes a variable number of int arguments.
Inside the function:
- Create a new instance of the Rectangle struct using &Rectangle{}.
- Extract values for width, height, x, and y from the variadic arguments.
- Create a new Base instance using NewBase and assign it to the embedded Base field of the Rectangle instance.
- Return a pointer to the newly created Rectangle struct.

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$  cat Main.go
```
package main

import (
	"fmt"

	"github.com/johnson-oragui/golang-almost-a-circle/models"
)

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
	TaskOneMain()
	
}
```

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go run Main.go

1

2

12

3

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ 

File: models/Rectangle.go, models/rectangle_test.go

## Task 3
**Validate fields**

Update the Rectangle struct in the models package with the following:

Setter Methods:

Implement setter methods for Width, Height, X, and Y.
Validate the input values:
If a value is not an integer, return an Error with the message "<field name> must be an integer".
If Width or Height is less than or equal to 0, return an Error with the message "<field name> must be > 0".
If X or Y is less than 0, return an Error with the message "<field name> must be >= 0".
NewRectangle Function:

Modify the NewRectangle function to validate the input values using the setter methods.
Return an error if any validation fails.

```
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
```

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go run Main.go

width: 10, height: 2, x: 0, y: 0, Id: 1

height must be > 0

width: 10, height: 2, x: 0, y: 0, Id: 12

width and height must be set and greather than zero

x must be >= 0

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ 

## Task 4
**Area First**

Update the struct Rectangle by adding the public method Area that returns the area value of the Rectangle instance.

```
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

```
johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go run Main.go

6

20

56

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ 

## Task 5
**Display**

Update the struct Rectangle by adding the public method Display that prints in stdout the Rectangle instance with the character # - you don’t need to handle x and y here.

```
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

```

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go run Main.go 

\# \# \# \#

\# \# \# \#

\# \# \# \#

\# \# \# \#

\# \# \# \#

\# \# \# \#

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ 

## Task 6
**String**
Update the struct Rectangle by overriding the String method so that it returns "[Rectangle] (<id>) <x>/<y> - <width>/<height>"
```
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


```
johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go run Main.go

[Rectangle] (12) 2/1 - 4/6

[Rectangle] (1) 1/0 - 5/5

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ 
