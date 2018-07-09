package show

import (
	"fmt"
)

// Print 使用fmt.Printf
func Print(s string, v interface{}) {
	fmt.Printf(s, v)
}
