package main

import (
	"fmt"
	"github.com/powerslider/ultimate-gopher/00-baby-gopher/02-packages/example/foo"
)

func main() {
	user := foo.NewUser("Homer", "Simpson", "s3cr37")
	// You can see the contents of the private information...
	fmt.Printf("%+v\n", user)
	// output:  {FirstName:Homer LastName:Simpson password:s3cr37}

	// You cannot access or change it directly....

	//fmt.Println(user.password)
	//user.password = "new"

	// output:
	//./main.go:16:18: user.password undefined (cannot refer to unexported field or method password)
	//./main.go:17:6: user.password undefined (cannot refer to unexported field or method password)
}

