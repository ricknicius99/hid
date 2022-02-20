package hid_test

import (
	"fmt"

	"github.com/karalabe/hid"
)

func ExampleSupported() {
	fmt.Println("Build configuration supported:", hid.Supported())
	// Output:
	// Build configuration supported: true
}
