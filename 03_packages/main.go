package main

import (
	"fmt"
	"math"

	"github.com/yahfiilham/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7)) // 2
	fmt.Println(math.Ceil(2.7)) // 3
	fmt.Println(math.Sqrt(49)) // 7
	fmt.Println(strutil.Reverse("olleh"))
}