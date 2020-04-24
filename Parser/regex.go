package Parser

import (
	log "github.com/sirupsen/logrus"
	"regexp"
)

// js example
// function a(bla){}
// eran todo

const ( // eran todo
	TODORegex     = ".*\\/\\/.*todo[[:space:],.`';!@#$%^&*=+:\\[\\]{}\\(\\)<>?]"
	GOFUNCRegex   = "func[[:space:]]*\\(?.*\\)?[[:space:]]*.*[[:space:]]*\\(.*\\)"
	JSFUNCREGEX   = "function"
	PYFUNCREGEX   = "def"
	OPEN_BRACKET  = "{"
	CLOSE_BRACKET = "}"
)

func funcRegex(extension string) string {
	switch extension {
	case "go":
		return GOFUNCRegex
	case "js":
		return JSFUNCREGEX
	case "py":
		return PYFUNCREGEX
	}
	return ""
}

func regex(text, regex string) int {
	r := regexp.MustCompile(regex)
	all := r.FindAllString(text, -1)
	if regex == TODORegex {
		log.Info("TODO ", text)
	}
	if all != nil {
		return len(all)
	}
	return 0
}
