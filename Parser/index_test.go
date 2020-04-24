package Parser

import (
	"testing"
)

func TestMainIndex(t *testing.T) {
	p := &Dir{}
	IndexFiles("", p)
	Parsetd(p)
	TODOsToMD(p)
}
