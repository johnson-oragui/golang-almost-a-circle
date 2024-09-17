### GOLANG-ALMOST-A-CIRCLE


### Tasks
0. **If it's not tested it doesn't work**
mandatory
All your files, classes and methods must be unit tested.

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go test
 ./models

ok      github.com/johnson-oragui/golang-almost-a-circle/models (cached)

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ 

File: models/models_test.go
   
1. **Base class**
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


func main() {

	b1 := models.NewBase()

	b2 := models.NewBase()

	b3 := models.NewBase()

	b4 := models.NewBase(12)

	b5 := models.NewBase()

	fmt.Printf("%d\n%d\n%d\n%d\n%d\n", b1.Id, b2.Id, b3.Id, b4.Id, b5.Id)
}
```

johnson1@DESKTOP-73V33K8:~/GOCODE/src/github.com/johnson-oragui/golang-almost-a-circle$ go run Main.go
1
2
3
12
4

Github Repo: golang-almost_a_circle
File: models/base.go
