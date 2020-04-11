package Parser

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

type AllTodos struct {
	td []Filetd
}

type Filetd struct {
	filename string
	td       []todo
}
type todo struct {
	isInFunc bool
	funcname string
	todo     string
}

func Parse(p string) (tds AllTodos) {
	if p == "" {
		p, _ = os.Getwd()
		log.Info(p)
		//todo: error handling
	}
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Error(err)
	}
	for _, f := range files {
		if f.IsDir() {
			if f.Name() == ".idea" || f.Name() == ".git" {
				continue
			}
			Parse(path.Join(p, f.Name()))
		} else {
			//log.Info("Parsing files: " , f.Name())
			parseFile(path.Join(p, f.Name()), &tds)
		}
	}
	return tds
}

func parseFile(p string, tds *AllTodos) {
	Ftd := Filetd{}
	f, err := os.Open(p)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s := scanner.Text()
		FindTODOs(s, Ftd)
		//todo: fix logic for correct func name

	}
	Ftd.filename = f.Name()
	tds.td = append(tds.td, Ftd)
}

func FindTODOs(line string, Ftd Filetd) bool {
	td := todo{}
	ok := Search(line)
	if !ok {
		return false
	}
	td.todo = line
	//todo isInFunc
	td.isInFunc = IsInFunc(line)
	td.funcname = GetFuncName(line)
	Ftd.td = append(Ftd.td, td)
	return true

	//TODO: Better search for TODOs - first run did not work becuase of it
}

func Search(line string) bool {
	line = strings.ToLower(line)
	if !strings.Contains(line, "//") {
		return false
	}
	//todo improve regex
	b, err := regexp.MatchString(".*todo.*", line)
	if err != nil {
		log.Error(err)
	}
	return b
}
