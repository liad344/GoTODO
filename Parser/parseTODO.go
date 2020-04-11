package Parser

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strings"
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
	f, err := os.Open(p)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		SaveToDOs(scanner.Text())
	}

}

func SaveToDOs(line string) {
	if !strings.Contains(line, "TODO") {
		return
	}
	if !strings.Contains(line, "TODO:") {
		return
	}
	if !strings.Contains(line, "//TODO:") {
		return
	}
	//TODO: Better search for TODOs

}
