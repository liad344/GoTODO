package Parser

import "strings"

func goParser(line string, td todo) {
	if strings.Contains(line, "func") {
		td.funcname = getFuncName(line)
	}

}

func getFuncName(line string) string {

}

func jsParser() {
	//...
}
