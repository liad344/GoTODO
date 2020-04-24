package Parser

type function struct {
	name  string
	index []int
}
type Dir struct {
	Name  string
	Files []*Indexed
	Dirs  []*Dir
}

type Indexed struct {
	Name      string
	Extension string
	// todo filed shoud remain unexported?
	len int
	f   []*function
	//i 		  []Interfaces?
	//c 		  []Classes?
	tds []*todo
}

type todo struct {
	isInFunc bool
	fn       *function
	todo     string
	index    int
}

var (
	Project Dir
)

func (d *Dir) abc() {

}

// js example
// function a(bla){}
// eran todo

const ( // eran todo
	TODORegex     = ".*\\/\\/.*todo[\\[:space:][](){}*&^%%$#@!~`?.,><:|;\"\\/\\\\]"
	GOFUNCRegex   = "func[[:space:]]*\\(?.*\\)?[[:space:]]*.*[[:space:]]*\\(.*\\)"
	JSFUNCREGEX   = "function"
	PYFUNCREGEX   = "def"
	OPEN_BRACKET  = "{"
	CLOSE_BRACKET = "}"
)
