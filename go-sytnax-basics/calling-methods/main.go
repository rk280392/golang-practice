package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.WriteRune('a')
	builder.WriteRune('b')
	builder.WriteRune('c')

	builder.WriteString("defgh")

	fmt.Println(builder.String())

}
