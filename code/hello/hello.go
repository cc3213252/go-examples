/**
@author: yudan.chen
@date: 2021/8/2
**/
package main

import (
	"example.com/user/hello/morestrings"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
