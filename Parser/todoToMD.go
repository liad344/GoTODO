package Parser

import (
	"github.com/romanyx/mdopen"
	log "github.com/sirupsen/logrus"
	"strings"
)

const SPACE = " "

func TODOsToMD(d *Dir) {
	md = ""
	f := strings.NewReader(todosToMD(d))
	opnr := mdopen.New()
	if err := opnr.Open(f); err != nil {
		log.Fatal(err)
	}
}

var md string

func todosToMD(d *Dir) string {
	for _, dir := range d.Dirs {
		if len(dir.Dirs) != 0 {
			todosToMD(&dir)
		}
		md += "**" + dir.Name + "**" + "<br/>"
		for _, file := range dir.Files {
			for _, td := range file.tds {
				if td.isInFunc {
					md += "FUNC" + td.fn.name + "FUNC" + "<br/>"
					log.Info("YoY", td.fn.name)
				}
				md += SPACE
				md += "```" + td.todo + "```" + "<br/>"
			}
		}
	}
	return md
}
