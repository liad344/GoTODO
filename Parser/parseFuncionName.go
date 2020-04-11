package Parser

import (
	"math/rand"
	"strconv"
)

func goParser(line string, td Filetd) {
	//if strings.Contains(line, "func") {
	//	td.funcname = getFuncName(line)
	//}
}

func GetFuncName(line string) string {
	i := rand.Int63n(3)
	return "VeryRealFunctionName" + strconv.Itoa(int(i))

}
func IsInFunc(line string) bool {
	return true
}

func jsParser() {
	//...
}
