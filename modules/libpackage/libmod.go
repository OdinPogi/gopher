package libpackage

import (
	"github.com/OdinPogi/gopher/modules/libpackage/abc"
	"github.com/OdinPogi/gopher/modules/libpackage/testing"
)

const Name = "Gopher Man"

func CallChild() string {
	return testing.TestIt()
}
func GetDEF() string {
	return abc.GetDEF()
}
