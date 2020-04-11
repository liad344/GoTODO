package Parser

import (
	"github.com/romanyx/mdopen"
	log "github.com/sirupsen/logrus"
	"strings"
)

func TODOsToMD(p string) {
	f := strings.NewReader("# Hello from markdown")

	opnr := mdopen.New()
	if err := opnr.Open(f); err != nil {
		log.Fatal(err)
	}
}
