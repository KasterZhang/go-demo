package decorator

import (
	"fmt"
)

func decorator(f func()) func(name string) {
	return func(name string) {
		fmt.Println("before decorator...", name)
		f()
		fmt.Println("after decorator...")
	}
}
