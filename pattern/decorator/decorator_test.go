package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	decorator(aloha)("Optional")
}

func aloha() {
	fmt.Println("aloha~~~")
}

func hello(name string) {
	fmt.Println("aloha~~~", name)
}
