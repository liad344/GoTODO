package Parser

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type todos struct {
	td []todo
}

type todo struct {
	line     []string
	filename string
	funcname string
	ok       bool
}

func parse(p string) (tds todos) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Error(err)
	}
	for _, f := range files {
		if f.IsDir() {
			parse(path.Join(p, f.Name()))
		} else {
			parseFile(path.Join(p, f.Name()), tds)
		}
	}
	return tds
}

func parseFile(p string, tds todos) {
	td := todo{}
	td.ok = true
	f, err := os.Open(p)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "func") {
			td.funcname = getFuncName(s)
			td.ok = false
		}
		//Assume togo is in current func
		//
		if FindTODOs(s, td) {
			if strings.Contains(s, "func") {
				td.funcname = getFuncName(s)
			}
		} else {
			td.ok = false
		}
	}
	td.filename = f.Name()
	tds.td = append(tds.td, td)
}

func FindTODOs(line string, td todo) bool {
	if !strings.Contains(line, "TODO") {
		return false
	}
	if !strings.Contains(line, "TODO:") {
		return false
	}
	if !strings.Contains(line, "//TODO:") {
		return false
	}
	td.line = append(td.line, line)
	return true

	//TODO: Better search for TODOs
}
