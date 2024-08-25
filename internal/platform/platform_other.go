//go:build windows || linux

package platform

import (
	"fmt"
)

func RemoveCodesigningSignature(path string) error {
	fmt.Println("Codesigning is a null op on your OS, TBD if this is OK.")
	return nil
}
