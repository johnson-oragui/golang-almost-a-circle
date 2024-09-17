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
File: models/base.go models/base_test.go

## Task 3
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

File: models/Rectangle.go models/rectangle_test.go
