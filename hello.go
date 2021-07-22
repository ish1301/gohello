package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"rocketsquare.com/site/gohello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

}
