package tolunacommons

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/mod/modfile"
)

func getModName() string {
	modcontent, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalln(err)
	}

	modulename := fmt.Sprintf("%s", modfile.ModulePath(modcontent))
	return string(modulename)
}
