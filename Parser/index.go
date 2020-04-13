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
func splitName(name string) (string, string) {
	s := strings.Split(name, ".")
	return s[0], s[1]
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

func parseFile(p string) (i Indexed) {
	f, err := os.Open(p)
	stat, _ := f.Stat()
	if err != nil {
		log.Error(err)
	}
	fnNum = 0
	if brackets != 0 {
		log.Error("Last file ended without closing bracket")
		brackets = 0
	}
	i.Name, i.Extension = splitName(stat.Name())
	name = i.Name
	i.f, i.len, i.tds = getFuncions(f, i.Extension)
	return i
}

var name string

func getFuncions(file *os.File, extension string) (f []function, len int, tds []todo) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	pf := funcRegex(extension)
	if pf == "" {
		for i := 1; scanner.Scan(); i++ {
			ln := scanner.Text()
			if regex(ln, TODORegex) > 0 {
				tds = append(tds, parsetodo(ln))
			}
		}
	} else {
		for i := 1; scanner.Scan(); i++ {
			ln := scanner.Text()
			if regex(ln, TODORegex) > 0 {
				tds = append(tds, parsetodo(ln))
			}
			parsefunc(ln, &f, pf, i)
		}
	}
	//todo it's ugly but then if/else for pf is checked once per file and not every line
	return f, len, tds
}

func parsetodo(ln string) (t todo) {
	t.todo = strings.ReplaceAll(ln, "//", "")
	return t
}

var brackets int
var fnNum int

func parsefunc(ln string, functions *[]function, functype string, i int) {
	f := function{
		name:  "",
		index: make([]int, 2),
	}

	if regex(ln, functype) > 0 {
		f.name = "!!!" + ln + "!!!"
		f.index[0] = i
		*functions = append(*functions, f)
		fnNum++ //Will be th num of the next function to come
		//brackets must be zero?
	}
	//if name == "root"{
	//	log.Info("brackets before " , brackets)
	//	log.Info("i " , i)
	//}
	lnHasBracket(ln)
	//if name == "root"{
	//	log.Info("brackets after " , brackets)
	//	log.Info(i , ") " , ln)
	//}
	if brackets == 0 && fnNum > 0 {
		(*functions)[fnNum-1].index[1] = i
	}

}

func lnHasBracket(ln string) {
	brackets += regex(ln, OPEN_BRACKET)
	brackets -= regex(ln, CLOSE_BRACKET)
}
func a() {
	go func() {

	}()
}
func (t todo) b() {
}

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
	if all != nil {
		return len(all)
	}
	return 0
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
