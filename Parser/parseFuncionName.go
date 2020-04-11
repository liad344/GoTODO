package Parser

import (
	log "github.com/sirupsen/logrus"
	"regexp"
)

func goParser(line string, td Filetd) {
	//if strings.Contains(line, "func") {
	//	td.funcname = getFuncName(line)
	//}
}

// dkd func dsk
func GetFuncName(line, function string) (name string) {
	r := function + "[[:space:]]"
	reg := regexp.MustCompile(r)
	loc := reg.FindStringIndex(line)
	if loc == nil {
		return ""
	}
	log.Info("name extraxted ", name, " from line ", line)
	return name

}
func IsInFunc(line string) bool {
	return true
}

func jsParser() {
	//...
}
