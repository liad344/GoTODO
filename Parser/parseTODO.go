package Parser

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

func parse(p string) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Error(err)
	}
	for _, f := range files {
		if f.IsDir() {
			parse(path.Join(p, f.Name()))
		} else {
			parseFile(path.Join(p, f.Name()))
		}
	}

}

func parseFile(p string) {

}
