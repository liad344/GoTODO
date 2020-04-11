package Parser

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func parse(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Error(err)
	}

}
