package Parser

import (
	"github.com/romanyx/mdopen"
	log "github.com/sirupsen/logrus"
	"strings"
)

func TODOsToMD() {
	f := strings.NewReader(todosToMD(*Tds))
	opnr := mdopen.New()
	if err := opnr.Open(f); err != nil {
		log.Fatal(err)
	}
}

func todosToMD(tds AllTodos) string {
	md := ""
	for _, td := range tds.td[0].td {
		if td.isInFunc {
			md += "**" + td.funcname + "**" + "<br/>"
		}
		md += "- " + td.todo + "<br/>"
	}
	return md
}
