package Parser

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var i int

func IndexFiles(p string, curdir *Dir) {
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
			IndexFiles(path.Join(p, f.Name()), &cd)
			curdir.Dirs = append(curdir.Dirs, &cd)
		} else {
			curdir.Files = append(curdir.Files, parseFile(path.Join(p, f.Name())))
		}
	}
	return
}
func Parsetd(d *Dir) {
	for _, dir := range d.Dirs {
		if len(dir.Dirs) != 0 {
			Parsetd(dir)
		}
		IndexTodos(dir)
	}
	IndexTodos(d)
}

func IndexTodos(dir *Dir) {
	for _, file := range dir.Files {
		for _, td := range file.tds {
			for _, fn := range file.f {
				if fn.index[0] <= td.index && fn.index[1] >= td.index {
					td.isInFunc = true
					td.fn = fn
				}
			}
		}
	}
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

func parseFile(p string) (i *Indexed) {
	i = &Indexed{}
	f, err := os.Open(p)
	stat, _ := f.Stat()
	if err != nil {
		log.Error(err)
	}
	fnNum = 0
	if brackets != 0 {
		log.Info(brackets)
		log.Info("Current file ", stat.Name())
		log.Error("Last file ended without closing bracket")
		brackets = 0
	}
	i.Name, i.Extension = splitName(stat.Name())
	name = i.Name
	i.f, i.len, i.tds = getFuncions(f, i.Extension)
	return i
}

var name string

func getFuncions(file *os.File, extension string) (f []*function, len int, tds []*todo) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	pf := funcRegex(extension)
	if pf == "" {
		for i := 1; scanner.Scan(); i++ {
			ln := scanner.Text()
			if regex(ln, TODORegex) > 0 {
				tds = append(tds, parsetodo(ln, i))
			}
		}
	} else {
		for i := 1; scanner.Scan(); i++ {
			ln := scanner.Text()
			if regex(ln, TODORegex) > 0 {
				tds = append(tds, parsetodo(ln, i))
			}
			parsefunc(ln, &f, pf, i)
		}
	}
	//todo it's ugly but then if/else for pf is checked once per file and not every line
	return f, len, tds
}

func parsetodo(ln string, i int) (t *todo) {
	t = &todo{}
	t.todo = parsetodoName(ln)
	t.isInFunc = false
	t.index = i
	return t
}
func parsetodoName(ln string) (name string) {
	return strings.ReplaceAll(ln[strings.Index(ln, "//"):], "//", "")
}

var brackets int
var fnNum int

func parsefunc(ln string, functions *[]*function, functype string, i int) {
	f := function{
		name:  "",
		index: make([]int, 2),
	}
	//todo makes an error if string has brackets {} and are non closing its counting it it
	if regex(ln, functype) > 0 {
		f.name = parsefuncName(ln)
		f.index[0] = i
		*functions = append(*functions, &f)
		fnNum++ //Will be th num of the next function to come
		//brackets must be zero?
	}
	//if name == "index"{
	//	log.Info("brackets before " , brackets)
	//	log.Info("i " , i)
	//}
	lnHasBracket(ln)
	//if name == "index"{
	//	log.Info("brackets after " , brackets)
	//	log.Info(i , ") " , ln)
	//}
	if brackets == 0 && fnNum > 0 {
		(*functions)[fnNum-1].index[1] = i
	}

}

func parsefuncName(ln string) (name string) {
	// We are getting something like this func[[:space:]]Name(args)
	// Or w/ receiver  func[[:space:]](r *receiver)Name(args)
	name = strings.Trim(ln, "func ")
	name = name[:strings.Index(name, "(")]
	//log.Info(name)
	return name
}

func lnHasBracket(ln string) {
	brackets += regex(ln, OPEN_BRACKET)
	brackets -= regex(ln, CLOSE_BRACKET)
}
