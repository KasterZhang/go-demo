package adaptor

import (
	"fmt"
	"testing"
)

func TestAdaptor(t *testing.T) {
	defer func() {
		fmt.Println("recovered ", recover())
	}()
	panic("sick")
}
