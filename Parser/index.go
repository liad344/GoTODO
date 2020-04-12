package Parser

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

var i int

func Parse(p string, curdir *Dir) {
	p, files := getFiles(p)
	//todo add project name
	for _, f := range files {
		if !include(f) {
			continue
		}
		if f.IsDir() {
			cd := Dir{
				Name: f.Name(),
			}
			Parse(path.Join(p, f.Name()), &cd)
			curdir.Dirs = append(curdir.Dirs, cd)
		} else {
			curdir.Files = append(curdir.Files, parseFile(path.Join(p, f.Name())))
		}
	}
	return
}
func getName(name string) (i Indexed) {
	s := strings.Split(name, ".")
	i.Name = s[0]
	i.Extension = s[1]
	// todo filed shoud remain unexported?
	return i
}

func getFiles(p string) (string, []os.FileInfo) {
	if p == "" {
		p, _ = os.Getwd()
		//todo: error handling
	}
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Error(err)
	}
	return p, files
}

func include(f os.FileInfo) bool {
	if f.Name() == ".idea" || f.Name() == ".git" || f.Name() == "GoTODO.exe" {
		return false
	}
	return true
	//todo ADD this to config
}

func parseFile(p string) (ix Indexed) {
	f, err := os.Open(p)
	if err != nil {
		log.Error(err)
	}
	stat, _ := f.Stat()
	ix = getName(stat.Name())
	return ix
}

func indexFile(f *os.File) Indexed {
	stat, _ := f.Stat()
	defer f.Close()
	return Indexed{Name: stat.Name()}
}

//
//func FindTODOs(line string, Ftd *Filetd) bool {
//	td := todo{}
//	ok := Search(line)
//	if !ok {
//		return false
//	}
//	td.todo = strings.ReplaceAll(line, "//", "")
//	//todo isInFunc
//	//td.isInFunc = IsInFunc(line)
//	//td.funcname = GetFuncName(line, "func")
//	//todo: change func for different programming languages like js = function
//	//Ftd.td = append(Ftd.td, td)
//	return true
//
//	//TODO: Better search for TODOs - first run did not work becuase of it
//}

func Search(line string) bool {
	line = strings.ToLower(line)
	if !strings.Contains(line, "//") {
		return false
	}
	b, _ := regexp.MatchString(TODORegex, line)
	return b
}
