package testdata

import (
	"log"

	_ "github.com/uudashr/go-module"
	"golang.org/x/mod/modfile"
	"gopkg.in/yaml.v2"
)

type Something struct{}

func aAllowedImport() {
	mfile, _ := modfile.Parse("go.mod", []byte{}, nil)

	log.Println(mfile)
}

func aBlockedImport() { // nolint: deadcode,unused
	data := []byte{}
	something := Something{}
	_ = yaml.Unmarshal(data, &something)

	log.Println(data)
}
