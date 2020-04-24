package Parser

import (
	_ "github.com/gomarkdown/markdown"
	_ "github.com/gomarkdown/markdown/parser"
	"github.com/romanyx/mdopen"
	_ "github.com/romanyx/mdopen"
	_ "github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	_ "io/ioutil"
	"strings"
	_ "strings"
)

const SPACE = " "

func TODOsToMD(d *Dir) {
	//md = ""
	//extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	//parser := parser.NewWithExtensions(extensions)
	//
	//md := []byte("[ ] yoo")
	//html := markdown.ToHTML(md, parser, nil)
	//_ = ioutil.WriteFile("todo.html", html, 0644)
	f := strings.NewReader(todosToMD(d))

	opnr := mdopen.New()
	opnr.Open(f)
	if err := opnr.Open(f); err != nil {
		log.Fatal(err)
	}
}

var md string

func todosToMD(d *Dir) string {
	for _, dir := range d.Dirs {
		if len(dir.Dirs) != 0 {
			todosToMD(dir)
		}
		md += "# " + dir.Name + "\n"
		for _, file := range dir.Files {
			md += "### " + "FILE: " + file.Name + "\n"
			if len(file.tds) == 0 {
				md += "```" + "	NO TODOs HERE" + "```" + "\n\n"
			}
			for _, td := range file.tds {
				if td.isInFunc {
					md += "##### " + "	FUNC " + td.fn.name + "\n\n"
				}
				md += SPACE
				md += "```" + "	TODO: " + td.todo + "```" + "\n\n"
			}
		}
	}
	return md
}
