package Parser

//
//import (
//	"github.com/romanyx/mdopen"
//	log "github.com/sirupsen/logrus"
//	"strings"
//)
//
//const SPACE = " "
//
//func TODOsToMD() {
//	f := strings.NewReader(todosToMD(*Tds))
//	opnr := mdopen.New()
//	if err := opnr.Open(f); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func todosToMD(tds AllTodos) string {
//	md := ""
//	lastFunc := ""
//	for _, ftd := range tds.td {
//		md += "**" + ftd.filename + "**" + "<br/>"
//		if len(ftd.td) == 0 {
//			md += SPACE + "No todos here <br/>"
//			continue
//		}
//		for _, td := range ftd.td {
//			if td.isInFunc && lastFunc != td.funcname {
//				lastFunc = td.funcname
//				md += "_" + td.funcname + "_" + "<br/>"
//			}
//			md += SPACE
//			md += "```" + td.todo + "```" + "<br/>"
//		}
//	}
//	return md
//}
