package Parser

import (
	"github.com/romanyx/mdopen"
	log "github.com/sirupsen/logrus"
	"strings"
)

func TODOsToMD(tds AllTodos) {
	f := strings.NewReader(todosToMD(tds))
	opnr := mdopen.New()
	if err := opnr.Open(f); err != nil {
		log.Fatal(err)
	}
}

func todosToMD(tds AllTodos) string {
	log.Info(tds)
	md := ""
	for _, Ftd := range tds.td {
		for _, td := range Ftd.td {
			if td.isInFunc {
				md += "#" + td.funcname
			}
			md += "-	" + td.todo + "\n"
		}
	}

	return md
}
