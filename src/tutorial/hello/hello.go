//executable commands must always use pakage main
package main

import (
		"fmt"
		"tutorial/stringutil"
)
func main() {
	fmt.Printf(stringutil.Reverse("hello world\n"))
}